<!-- ![](media/logo_med.png) -->

# 🎭 Preludio

### A PRQL based data transformation language

Preludio is a data transformation language based on PRQL. It is a language that allows you to transform and
manipulate data in a simple and intuitive way, batteries included.

No libraries or external dependencies are required to run the language.

### Examples

Read and clean up a CSV file, then store the result in a variable called `clean`:

```
clean := (
  rcsv! p'test_files/Cars.csv' del:';' head:true
  strRepl! [MPG, Displacement, Horsepower, Acceleration] old:',' new:'.'
  asFlt! [MPG, Displacement, Horsepower, Acceleration]
  sort! [-Origin, Cylinders, -MPG]
)
```

```
europe5Cylinders := (
  from! clean
  filter! Cylinders == 5 and Origin == 'Europe'
)
```

Derive new columns and write the result to a CSV file:

```
from! clean
derive! [
  Stat = ((MPG * Cylinders * Displacement) / Horsepower * Acceleration) / Weight,
  CarOrigin = Car + ' - ' + Origin
]
filter! Stat > 1.3
select! [Car, Origin, Stat]
wcsv! p'test_files/Cars1.csv' del: '\t'
```

Create a new table by joining two tables:

```
continents := (
  new! [
    Continent = ['Asia', 'America', 'Europe'],
    Origin = ['Japan', 'US', 'Europe']
  ]
)

joined := (
  from! clean
  leftj! continents on: [Origin]
  select! [Car, Origin, Continent]
  sort! [Continent, Origin]
)
```

![](media/repl_example.gif)

### Data Types

The language supports the following data types:

- `boolean` ie: `true`/`false`
- `integer` ie: `1`, `2`, `3`
- `range` ie: `1..10`, `1..10:2`
- `float` ie: `1.0`, `2e-3`, `4.5E+6`
- `string` which has some variants:
  - plain `'hello world'`, `"foo bar"`
  - raw `r'\d'`, `r"\a"`
  - path `p'c:\temp'`, `p"/home/user"`
- `regex` ie: `x'he(l){2}o'`, `x"f(o){2} bar"`
- `date` ie: `d'2021-08-20'`, `d"2021-08-20"`
- `duration` ie: `1:h`, `2:milliseconds`, `3:us`

In addition, the language supports the following data structures:

- `list` ie: `a := [1, 2, 3]`, `b := [i = [1, 2], f = [3.0, 4.0]]`

### Operators

The language supports the following operators:

- `+` (addition)
- `-` (subtraction)
- `*` (multiplication)
- `/` (division)
- `%` (modulo)
- `^` (exponentiation)
- `and` (logical and)
- `or` (logical or)
- `not` (logical not)
- `==` (equality)
- `!=` (inequality)
- `>` (greater than)
- `>=` (greater than or equal to)
- `<` (less than)
- `<=` (less than or equal to)

- string interpolation `f'I have {1 + 2} apples'`

### Built-in Functions

The language supports the following built-in functions:

- `from` initializes a pipeline, ie: `from table`
- `new` creates a new dataframe, ie: `new [a = [1, 2], b = [3, 4]]`
- `rcsv` reads a CSV file, ie: `rcsv p'c:\temp\file.csv' del:',' head:true`
- `wcsv` writes a CSV file ie: `wcsv p'c:\temp\file.csv' del:','`
- `filter` filters rows, ie: `filter a > 1`
- `select` selects columns, ie: `select [a, b]`
- `sort` sorts rows, ie: `sort [a, -b]`
- `derive` adds new columns from the existing ones, ie: `derive [c = a + b]`

### Features

- [x] Arithmetic and logical operators
- [x] Read and write CSV files
- [x] Derive new columns
- [x] Select columns
- [x] Filter rows
- [x] Sort rows
- [x] Join tables
- [ ] Group by and aggregate

### Installation

To run it, you need to have [Go](https://golang.org/doc/install) installed.
Once you have Go, you can clone this repository.

To run the program, you can use the following command:

```bash
go mod tidy
go run .
```

### Future Features

- [x] Move to [Gandalff](https://github.com/caerbannogwhite/preludio/tree/main/core/gandalff) library
- [ ] Add date/time data types
- [ ] Add statistical functions
- [ ] Add support for Excel files
- [ ] Add support for XPT files
- [ ] VS Code extension
- [ ] Add support for SAS7BDAT files
- [ ] Add support for SPSS files
- [ ] Database connections (SQL, MongoDB, etc.)

In case the language becomes quite successful, I will consider adding:

- [ ] Text editor/Ide (in browser and app)
- [ ] Plots (https://github.com/gonum/plot)
- [ ] Interactive plots and tables
- [ ] Integration with other languages (Python, R, etc.)
- [ ] Integration with other tools (Jupyter, etc.)
- [ ] Integration with OpenAI (https://openai.com/blog/openai-api/), ie. image to table

### Contributing

If you want to contribute to this project, you can do so by forking the repository and submitting a pull request.

### Developers

If the grammar is changed, the parser must be regenerated. To do this, run the following command:

(on Windows)

```
make.ps1
```

### New Ideas

- Add datetime series to Gandalff
- List can be indexed with integers, ranges, strings and regex

### Log

- **20 / 08 / 2023** After exactly one year from the first commit, Preludio is fairly stable and usable. The language is still missing a few core features (like `join` and aggregators, already supported by Gandalff), but it is already possible to perform many operations with it.
- **02 / 08 / 2023** Preludio is now using the Gandalff library for managing data.
- **21 / 03 / 2023** First publishing of the repository. Many things are still not working.
- **18 / 03 / 2023** Gandalff library: fist commit.
- **20 / 08 / 2022** Preludio: fist commit.
