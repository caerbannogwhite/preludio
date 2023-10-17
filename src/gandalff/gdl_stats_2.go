package gandalff

import (
	"fmt"
	"math"
	"sort"
)

func _Mean(a []float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum / float64(len(a))
}

func StdDev(a []float64) float64 {
	_mean := _Mean(a)
	var sum float64
	for _, v := range a {
		sum += (v - _mean) * (v - _mean)
	}
	return math.Sqrt(sum / float64(len(a)))
}

// This functions computes the t-tests.
func TTest(a, b []float64) (t, p float64) {
	// Compute the _means.
	_meanA := _Mean(a)
	_meanB := _Mean(b)
	// Compute the standard deviations.
	sdA := StdDev(a)
	sdB := StdDev(b)
	// Compute the t-test.
	t = (_meanA - _meanB) / math.Sqrt((sdA*sdA/float64(len(a)))+(sdB*sdB/float64(len(b))))
	// Compute the p-value.
	p = 2 * (1 - StudentT(float64(len(a)+len(b)-2), t))
	return
}

// This function computes the Student's t-distribution.
func StudentT(nu float64, x float64) float64 {
	// Compute the first part of the function.
	a := math.Gamma((nu + 1) / 2)
	// Compute the second part of the function.
	b := math.Gamma(nu / 2)
	// Compute the third part of the function.
	c := math.Pow(nu*math.Pi, 0.5)
	// Compute the fourth part of the function.
	d := math.Pow(1+(x*x/nu), -(nu+1)/2)
	// Return the result.
	return a / (b * c * d)
}

// This function computes the inverse of the Student's t-distribution.
func InvStudentT(nu float64, p float64) float64 {
	// Compute the first part of the function.
	a := math.Pow(nu*math.Pi, 0.5)
	// Compute the second part of the function.
	b := math.Pow(1-p, 1/nu)
	// Compute the third part of the function.
	c := math.Pow(1-p, 1/nu)
	// Compute the fourth part of the function.
	d := math.Pow(1-p, 1/nu)
	// Return the result.
	return a * (b - c) / (1 + d)
}

// Compute the confidence interval.
func ConfidenceInterval(a []float64, p float64) (float64, float64) {
	// Compute the _mean.
	_mean := _Mean(a)
	// Compute the standard deviation.
	sd := StdDev(a)
	// Compute the t-value.
	t := InvStudentT(float64(len(a)-1), p)
	// Compute the confidence interval.
	ci := t * (sd / math.Sqrt(float64(len(a))))
	// Return the result.
	return _mean - ci, _mean + ci
}

// This function computes the p-value for the t-test.
func TTestPValue(a, b []float64) float64 {
	// Compute the t-test.
	t, _ := TTest(a, b)
	// Compute the p-value.
	return 2 * (1 - StudentT(float64(len(a)+len(b)-2), t))
}

// Compute the z-score.
func ZScore(a []float64) float64 {
	// Compute the _mean.
	_mean := _Mean(a)
	// Compute the standard deviation.
	sd := StdDev(a)
	// Compute the z-score.
	return (_mean - 0) / sd
}

// This function computes the p-value for the z-test.
func ZTestPValue(a []float64) float64 {
	// Compute the z-score.
	z := ZScore(a)
	// Compute the p-value.
	return 2 * (1 - Normal(z))
}

// This function computes the normal distribution.
func Normal(x float64) float64 {
	// Compute the first part of the function.
	a := 1 / math.Sqrt(2*math.Pi)
	// Compute the second part of the function.
	b := math.Exp(-x * x / 2)
	// Return the result.
	return a * b
}

// This function computes the inverse of the normal distribution.
func InvNormal(p float64) float64 {
	// Compute the first part of the function.
	a := math.Sqrt(2 * math.Pi)
	// Compute the second part of the function.
	b := math.Pow(math.E, -p*p/2)
	// Return the result.
	return a * b
}

// This function computes the confidence interval.
func ZConfidenceInterval(a []float64, p float64) (float64, float64) {
	// Compute the _mean.
	_mean := _Mean(a)
	// Compute the standard deviation.
	sd := StdDev(a)
	// Compute the z-value.
	z := InvNormal(p)
	// Compute the confidence interval.
	ci := z * (sd / math.Sqrt(float64(len(a))))
	// Return the result.
	return _mean - ci, _mean + ci
}

// Simple linear regression.
func SimpleLinearRegression(x, y []float64) (float64, float64) {
	// Compute the _means.
	_meanX := _Mean(x)
	_meanY := _Mean(y)
	// Compute the standard deviations.
	sdX := StdDev(x)
	sdY := StdDev(y)
	// Compute the correlation coefficient.
	r := Correlation(x, y)
	// Compute the slope.
	slope := r * (sdY / sdX)
	// Compute the intercept.
	intercept := _meanY - (slope * _meanX)
	// Return the result.
	return slope, intercept
}

// This function computes the correlation coefficient.
func Correlation(x, y []float64) float64 {
	// Compute the _means.
	_meanX := _Mean(x)
	_meanY := _Mean(y)
	// Compute the standard deviations.
	sdX := StdDev(x)
	sdY := StdDev(y)
	// Compute the correlation coefficient.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += (x[i] - _meanX) * (y[i] - _meanY)
	}
	return sum / (float64(len(x)) * sdX * sdY)
}

// This function computes the p-value for the correlation coefficient.
func CorrelationPValue(x, y []float64) float64 {
	// Compute the correlation coefficient.
	r := Correlation(x, y)
	// Compute the p-value.
	return 2 * (1 - StudentT(float64(len(x)-2), math.Abs(r)))
}

// This function computes the p-value for the simple linear regression.
func SimpleLinearRegressionPValue(x, y []float64) float64 {
	// Compute the correlation coefficient.
	r := Correlation(x, y)
	// Compute the p-value.
	return 2 * (1 - StudentT(float64(len(x)-2), math.Abs(r)))
}

// This function computes the p-value for the multiple linear regression.
func MultipleLinearRegressionPValue(x [][]float64, y []float64) float64 {
	// Compute the correlation coefficient.
	r := MultipleLinearRegression(x, y)
	// Compute the p-value.
	return 2 * (1 - StudentT(float64(len(x)-2), math.Abs(r)))
}

// Multiple linear regression.
func MultipleLinearRegression(x [][]float64, y []float64) float64 {
	// Compute the standard deviations.
	sdY := StdDev(y)
	// Compute the correlation coefficient.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += Correlation(x[i], y) * (StdDev(x[i]) / sdY)
	}
	return sum / float64(len(x))
}

// This function computes the p-value for the chi-square test.
func ChiSquareTestPValue(a [][]float64) float64 {
	// Compute the chi-square test.
	chi, _ := ChiSquareTest(a)
	// Compute the p-value.
	return 1 - ChiSquare(float64(len(a)-1), chi)
}

// This function computes the chi-square test.
func ChiSquareTest(a [][]float64) (float64, float64) {
	// Compute the row sums.
	var rowSum []float64
	for i := 0; i < len(a); i++ {
		var sum float64
		for j := 0; j < len(a[i]); j++ {
			sum += a[i][j]
		}
		rowSum = append(rowSum, sum)
	}
	// Compute the column sums.
	var colSum []float64
	for j := 0; j < len(a[0]); j++ {
		var sum float64
		for i := 0; i < len(a); i++ {
			sum += a[i][j]
		}
		colSum = append(colSum, sum)
	}
	// Compute the total sum.
	var totalSum float64
	for i := 0; i < len(rowSum); i++ {
		totalSum += rowSum[i]
	}
	// Compute the chi-square test.
	var chi float64
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			chi += math.Pow(a[i][j]-((rowSum[i]*colSum[j])/totalSum), 2) / ((rowSum[i] * colSum[j]) / totalSum)
		}
	}
	// Compute the degrees of freedom.
	df := float64((len(a) - 1) * (len(a[0]) - 1))
	// Compute the p-value.
	p := 1 - ChiSquare(df, chi)
	// Return the result.
	return chi, p
}

// This function computes the chi-square distribution.
func ChiSquare(df float64, x float64) float64 {
	// Compute the first part of the function.
	a := math.Pow(x, df/2)
	// Compute the second part of the function.
	b := math.Pow(2, df/2)
	// Compute the third part of the function.
	c := math.Gamma(df / 2)
	// Compute the fourth part of the function.
	d := math.Pow(math.E, -x/2)
	// Return the result.
	return a / (b * c * d)
}

// This function computes the inverse chi-square distribution.
func InvChiSquare(df float64, p float64) float64 {
	// Compute the first part of the function.
	a := math.Pow(2, df/2)
	// Compute the second part of the function.
	b := math.Gamma(df / 2)
	// Compute the third part of the function.
	c := math.Pow(p, -2/df)
	// Return the result.
	return a * b * c
}

// This function computes the F distribution.
func FDistribution(df1 float64, df2 float64, x float64) float64 {
	// Compute the first part of the function.
	a := math.Pow(df1*x, df1)
	// Compute the second part of the function.
	b := math.Pow(df2, df2)
	// Compute the third part of the function.
	c := math.Pow(df1*x+df2, df1+df2)
	// Compute the fourth part of the function.
	d := Beta(df1/2, df2/2)
	// Return the result.
	return (a * b) / (c * d)
}

// This function computes the inverse F distribution.
func InvFDistribution(df1 float64, df2 float64, p float64) float64 {
	// Compute the first part of the function.
	a := BetaInv(p, df1/2, df2/2)
	// Compute the second part of the function.
	b := math.Pow(df2/df1, df2/2)
	// Compute the third part of the function.
	c := math.Pow(df1/df2, df1/2)
	// Return the result.
	return (a * b) / c
}

// This function computes the beta distribution.
func Beta(a float64, b float64) float64 {
	// Compute the first part of the function.
	a1 := math.Gamma(a)
	// Compute the second part of the function.
	b1 := math.Gamma(b)
	// Compute the third part of the function.
	c := math.Gamma(a + b)
	// Return the result.
	return a1 * b1 / c
}

// This function computes the inverse beta distribution.
func BetaInv(p float64, a float64, b float64) float64 {
	// Compute the first part of the function.
	a1 := math.Gamma(a)
	// Compute the second part of the function.
	b1 := math.Gamma(b)
	// Compute the third part of the function.
	c := math.Gamma(a + b)
	// Compute the fourth part of the function.
	d := math.Pow(p, 1/a)
	// Compute the fifth part of the function.
	e := math.Pow(1-p, 1/b)
	// Return the result.
	return (a1 * b1 * d) / (c * e)
}

// Anova performs a one-way analysis of variance.
func Anova(x [][]float64, y []float64) (float64, float64) {
	// Compute the total sum of squares.
	tss := TotalSumOfSquares(y)
	// Compute the residual sum of squares.
	rss := ResidualSumOfSquares(x, y)
	// Compute the _mean sum of squares.
	mss := tss - rss
	// Compute the degrees of freedom.
	df1 := float64(len(x) - 1)
	df2 := float64(len(y) - len(x))
	// Compute the F statistic.
	f := mss / rss
	// Compute the p-value.
	p := 1 - FDistribution(df1, df2, f)
	// Return the result.
	return f, p
}

// This function computes the total sum of squares.
func TotalSumOfSquares(y []float64) float64 {
	// Compute the _mean of the data.
	_mean := _Mean(y)
	// Compute the sum of squares.
	var sum float64
	for i := 0; i < len(y); i++ {
		sum += math.Pow(y[i]-_mean, 2)
	}
	// Return the result.
	return sum
}

// This function computes the residual sum of squares.
func ResidualSumOfSquares(x [][]float64, y []float64) float64 {
	// Compute the sum of squares.
	var sum float64
	for i := 0; i < len(x); i++ {
		// Compute the _mean of the data.
		_mean := _Mean(x[i])
		// Compute the sum of squares.
		for j := 0; j < len(x[i]); j++ {
			sum += math.Pow(x[i][j]-_mean, 2)
		}
	}
	// Return the result.
	return sum
}

// This function computes the binomial distribution.
func Binomial(n int, p float64, x int) float64 {
	// Compute the first part of the function.
	a := math.Pow(p, float64(x))
	// Compute the second part of the function.
	b := math.Pow(1-p, float64(n-x))
	// Compute the third part of the function.
	c := Combination(n, x)
	// Return the result.
	return a * b * c
}

// This function computes the inverse binomial distribution.
func InvBinomial(n int, p float64, pval float64) int {
	// Compute the first part of the function.
	a := math.Pow(1-p, float64(n))
	// Compute the second part of the function.
	b := math.Pow(p, float64(n))
	// Compute the third part of the function.
	c := math.Pow(1-pval, 1/float64(n))
	// Return the result.
	return int(math.Log(b/a) / math.Log(c))
}

// This function computes the combination.
func Combination(n int, x int) float64 {
	// Compute the first part of the function.
	a := Factorial(n)
	// Compute the second part of the function.
	b := Factorial(x)
	// Compute the third part of the function.
	c := Factorial(n - x)
	// Return the result.
	return float64(a) / float64(b*c)
}

// This function computes the factorial.
func Factorial(n int) int {
	// Compute the factorial.
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	// Return the result.
	return result
}

// This function prints an histogram.
func PrintHistogram(x []float64, n int) {
	// Compute the minimum and maximum values.
	min, max := MinMax(x)
	// Compute the bin width.
	width := (max - min) / float64(n)
	// Compute the bins.
	bins := make([]int, n)
	for i := 0; i < len(x); i++ {
		// Compute the index of the bin.
		index := int((x[i] - min) / width)
		// Increment the bin.
		bins[index]++
	}
	// Print the histogram.
	for i := 0; i < n; i++ {
		// Compute the lower bound.
		lower := min + float64(i)*width
		// Compute the upper bound.
		upper := min + float64(i+1)*width
		// Print the bin.
		fmt.Printf("%f - %f: %d\n", lower, upper, bins[i])
	}
}

// This function returns the minimum and maximum values.
func MinMax(x []float64) (float64, float64) {
	// Compute the minimum and maximum values.
	min := x[0]
	max := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
		if x[i] > max {
			max = x[i]
		}
	}
	// Return the result.
	return min, max
}

// This function computes the percentile.
func Percentile(x []float64, p float64) float64 {
	// Compute the index.
	index := p * float64(len(x))
	// Compute the percentile.
	return x[int(index)]
}

// This function computes the _median.
func _Median(x []float64) float64 {
	// Compute the _median.
	return Percentile(x, 0.5)
}

// This function computes the mode.
func Mode(x []float64) float64 {
	// Compute the mode.
	var mode float64
	var count int
	for i := 0; i < len(x); i++ {
		// Compute the number of occurrences.
		var temp int
		for j := 0; j < len(x); j++ {
			if x[i] == x[j] {
				temp++
			}
		}
		// Update the mode.
		if temp > count {
			mode = x[i]
			count = temp
		}
	}
	// Return the result.
	return mode
}

// This function computes the variance.
func Variance(x []float64) float64 {
	// Compute the _mean.
	_mean := _Mean(x)
	// Compute the variance.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += math.Pow(x[i]-_mean, 2)
	}
	// Return the result.
	return sum / float64(len(x))
}

// This function computes the r squared.
func RSquared(x []float64, y []float64) float64 {
	// Compute the r squared.
	return 1 - Variance(x)/Variance(y)
}

// This function computes the covariance.
func Covariance(x []float64, y []float64) float64 {
	// Compute the covariance.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += (x[i] - _Mean(x)) * (y[i] - _Mean(y))
	}
	// Return the result.
	return sum / float64(len(x))
}

// This function computes a kernel density estimate.
func KernelDensityEstimate(x []float64, h float64) []float64 {
	// Compute the kernel density estimate.
	var y []float64
	for i := 0; i < len(x); i++ {
		// Compute the first part of the function.
		a := 1 / (math.Sqrt(2*math.Pi) * h)
		// Compute the second part of the function.
		b := math.Exp(-math.Pow(x[i], 2) / (2 * math.Pow(h, 2)))
		// Append the result.
		y = append(y, a*b)
	}
	// Return the result.
	return y
}

// This function computes the k-nearest neighbors.
func KNearestNeighbors(x []float64, y []float64, k int) [][]float64 {
	// Compute the k-nearest neighbors.
	var neighbors [][]float64
	for i := 0; i < len(x); i++ {
		// Compute the distances.
		var distances []float64
		for j := 0; j < len(x); j++ {
			distances = append(distances, math.Sqrt(math.Pow(x[i]-x[j], 2)+math.Pow(y[i]-y[j], 2)))
		}
		// Sort the distances.
		sort.Float64s(distances)
		// Compute the k-nearest neighbors.
		var temp []float64
		for j := 0; j < k; j++ {
			temp = append(temp, distances[j])
		}
		neighbors = append(neighbors, temp)
	}
	// Return the result.
	return neighbors
}

// This function computes the k-_means clustering.
func K_MeansClustering(x []float64, y []float64, k int) [][]float64 {
	// Compute the k-_means clustering.
	var clusters [][]float64
	for i := 0; i < k; i++ {
		// Compute the clusters.
		var temp []float64
		for j := 0; j < len(x); j++ {
			if y[j] == float64(i) {
				temp = append(temp, x[j])
			}
		}
		clusters = append(clusters, temp)
	}
	// Return the result.
	return clusters
}

// This function computes the determinant.
func Determinant(x [][]float64) float64 {
	// Compute the determinant.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += x[0][i] * Cofactor(x, 0, i)
	}
	// Return the result.
	return sum
}

// This function computes the cofactor.
func Cofactor(x [][]float64, i int, j int) float64 {
	// Compute the cofactor.
	return math.Pow(-1, float64(i+j)) * Minor(x, i, j)
}

// This function computes the minor.
func Minor(x [][]float64, i int, j int) float64 {
	// Compute the minor.
	var temp [][]float64
	for k := 0; k < len(x); k++ {
		if k != i {
			var temp2 []float64
			for l := 0; l < len(x); l++ {
				if l != j {
					temp2 = append(temp2, x[k][l])
				}
			}
			temp = append(temp, temp2)
		}
	}
	// Return the result.
	return Determinant(temp)
}

// This function computes the inverse.
func Inverse(x [][]float64) [][]float64 {
	// Compute the inverse.
	var temp [][]float64
	for i := 0; i < len(x); i++ {
		var temp2 []float64
		for j := 0; j < len(x); j++ {
			temp2 = append(temp2, Cofactor(x, i, j)/Determinant(x))
		}
		temp = append(temp, temp2)
	}
	// Return the result.
	return Transpose(temp)
}

// This function computes the transpose.
func Transpose(x [][]float64) [][]float64 {
	// Compute the transpose.
	var temp [][]float64
	for i := 0; i < len(x); i++ {
		var temp2 []float64
		for j := 0; j < len(x); j++ {
			temp2 = append(temp2, x[j][i])
		}
		temp = append(temp, temp2)
	}
	// Return the result.
	return temp
}

// This function computes the dot product.
func DotProduct(x []float64, y []float64) float64 {
	// Compute the dot product.
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += x[i] * y[i]
	}
	// Return the result.
	return sum
}

// This function computes the subtract.
func Subtract(x []float64, y []float64) []float64 {
	// Compute the subtract.
	var temp []float64
	for i := 0; i < len(x); i++ {
		temp = append(temp, x[i]-y[i])
	}
	// Return the result.
	return temp
}

// This function computes the multiply.
func Multiply(x [][]float64, y [][]float64) [][]float64 {
	// Compute the multiply.
	var temp [][]float64
	for i := 0; i < len(x); i++ {
		var temp2 []float64
		for j := 0; j < len(x); j++ {
			temp2 = append(temp2, x[i][j]*y[i][j])
		}
		temp = append(temp, temp2)
	}
	// Return the result.
	return temp
}

// This function computes the add.
func Add(x []float64, y []float64) []float64 {
	// Compute the add.
	var temp []float64
	for i := 0; i < len(x); i++ {
		temp = append(temp, x[i]+y[i])
	}
	// Return the result.
	return temp
}

// This function computes the divide.
func Divide(x []float64, y []float64) []float64 {
	// Compute the divide.
	var temp []float64
	for i := 0; i < len(x); i++ {
		temp = append(temp, x[i]/y[i])
	}
	// Return the result.
	return temp
}
