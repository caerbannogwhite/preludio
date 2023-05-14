### Introduction

I originally started to build Gandalff because I noticed a lack of good and complete data analysis tools for Go. However, the main purpose was to provide a performant backbone library for my former project [Preludio](), a programming language for data transformation based on PRQL.
However, I ended up enjoying the development of Gandalff a bit too much and now I'm trying to make it as complete as possible and as performant as Polars.

One of the most significant operations for a data analysis library is the `groupby` function.
By reading the Polars documentation I noticed this quite fun [challenge](https://h2oai.github.io/db-benchmark/): it's a benchmarking of the `groupby` and `join` functions for the most popular data analysis libraries on 3 datasets of 0.5 GB, 5 GB, and 50 GB size.
Not being interested yet in the memory usage of Gandalff and seeing if it can survive the 50 GB dataset, I decided to use some smaller datasets and see how Gandalff performs.

### Baseline

I first generated 4 datasets of 0.5 MB, 5 MB, 50 MB, and 500 MB using the [R script](https://github.com/h2oai/db-benchmark/blob/master/_data/groupby-datagen.R) provided in the challenge. These datasets have respectively 10'000, 100'000, 1'000'000, and 10'000'000 rows and 9 columns (3 strings, 5 integers, and 1 float).

Here is what the first dataset looks like:

|  id1  |  id2  |     id3      | id4 | id5 | id6 | v1  | v2  |    v3     |
| :---: | :---: | :----------: | :-: | :-: | :-: | :-: | :-: | :-------: |
| id016 | id056 | id0000000042 | 82  | 64  | 50  |  4  |  3  | 26.407777 |
| id039 | id075 | id0000000067 | 30  | 28  | 17  |  1  |  6  | 7.913725  |
| id047 | id077 | id0000000027 | 40  | 67  | 23  |  5  | 14  | 54.691464 |
| id043 | id072 | id0000000055 | 92  |  6  | 21  |  1  |  2  | 13.573742 |

I then collected the execution time for the first 5 questions (out of 10) of the `groupby` task for each dataset and each library.
The questions are the following:

- Q1: `sum v1 by id1`
- Q2: `sum v1 by id1:id2`
- Q3: `sum v1 mean v3 by id3`
- Q4: `mean v1:v3 by id4`
- Q5: `sum v1:v3 by id6`

I skipped the remaining 5 questions because they are not supported by Gandalff yet.

pandas version: 1.3.5
polars version: 0.17.10

Looking at the ratios between Gandalff and Polars on the average execution time over all the questions for a given input size, we can see that Gandalff is almost _8x_ slower than Polars on the 500 MB dataset (worst case scenario).

|    Input Size |   Gandalff |   Polars |   Ratio |
|--------------:|-----------:|---------:|--------:|
|     10000     |      1.036 |    0.414 |     2.5 |
|    100000     |      5.454 |    0.852 |     6.4 |
|         1e+06 |     64.49  |   27.992 |     2.3 |
|         1e+07 |   1063.94  |  136.322 |     7.8 |

### Gandalff's First Attempt

My first naive implementation of the `groupby` function was to use a `map[type][]int` to store the indices of the unique values in a series.
I suddenly realized thet for the following step, which was: implementing the `groupby` function for multiple series, in a dataframe, this approach required a slightly different data structure, which is `[]map[type][]int`. In this way, I could store the indices of the rows that have the same values for the first series, then for the second series, and so on.

The dataframe `GroupBy` algorithm works more or less in this way:

- it first calls the `Group` function for the first series in the set of the grouping series.
  This generates a first partition of the rows, which is the `map[type][]int` that I described above.
- then, it iterates the remaining series, using another method (that I called `SubGroup`) that takes the current partition of the rows and generates a new one by grouping the values of the current series.

The `SubGroup` for a series of `int` values can be represented by the following code:

```go
nullGroups = make([][]int, partitions.GetGroupsCount())

for gi, g := range indices {

    // initialize embedded partitions
    embeddedPartitions[gi] = make(map[int][]int)
    nullGroups[gi] = make([]int, 0)

    for _, idx := range g {
        if s.IsNull(idx) {
            nullGroups[gi] = append(nullGroups[gi], idx)
        } else {
            if embeddedPartitions[gi][s.data[idx]] == nil {
                embeddedPartitions[gi][s.data[idx]] = make([]int, 0)
            }
            embeddedPartitions[gi][s.data[idx]] = append(embeddedPartitions[gi][s.data[idx]], idx)
        }
    }
}
```

The `embeddedPartitions` is the _sub-partitio_, or the _partition_ nested in the current one.

Let's consider the following example:

| Name | Department | Age | Salary |
| :--: | :--------: | :-: | :----: |
|  A   |     HR     | 30  |  5000  |
|  B   |     HR     | 40  |  6000  |
|  C   |     IT     | 25  |  4000  |
|  D   |     IT     | 30  |  5000  |
|  E   |     IT     | 35  |  5500  |

### First steps

From the flamegraph, I noticed that for Q1, Q3, Q4, and Q5 the test code spends a bit more time in the aggregation function than in the `groupby` function.
For Q2, however, the string sub-grouping function is a time sink.

Checking the code for the sum aggregation function, I noticed a quite embarrassing mistake: there is no need to check if an element is null or not, because the `groupby` function already takes care of that and all the null values (if any) are stored in a separate group.
An if statement inside a loop that has to run for millions of times is a big no-no.
Also, calling `series.Get(i)` for each element is not a good idea.

```go
sum := make([]float64, len(groups))
switch series := s.(type) {
// ...
case GDLSeriesFloat64:
    if series.isNullable {
        for gi, group := range groups {
            for _, i := range group {
                if !series.IsNull(i) {
                    sum[gi] += series.Get(i).(float64)
                }
            }
        }
        return sum
    } else {
        for gi, group := range groups {
            for _, i := range group {
                sum[gi] += series.Get(i).(float64)
            }
        }
        return sum
    }
// ...
}
```

The new code looks like this:

```go
sum := make([]float64, len(groups))
switch series := s.(type) {
// ...
case GDLSeriesFloat64:
    data := *series.__getDataPtr()
    for gi, group := range groups {
        for _, i := range group {
            sum[gi] += data[i]
        }
    }
    return sum
// ...
}
```

The results are reported in the following table:

|    Input Size |   Gandalff 1 |   Polars |   Ratio |
|--------------:|-------------:|---------:|--------:|
|     10000     |        0.64  |    0.414 |    1.54 |
|    100000     |        4.036 |    0.852 |    4.73 |
|         1e+06 |       44.248 |   27.992 |    1.58 |
|         1e+07 |      685.986 |  136.322 |    5.03 |

[](https://www.cockroachlabs.com/blog/vectorized-hash-joiner/)
