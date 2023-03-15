package stats

import "math"

// This functions computes the t-tests.
func TTest(a, b []float64) (t, p float64) {
	// Compute the means.
	meanA := Mean(a)
	meanB := Mean(b)
	// Compute the standard deviations.
	sdA := StdDev(a)
	sdB := StdDev(b)
	// Compute the t-test.
	t = (meanA - meanB) / math.Sqrt((sdA*sdA/float64(len(a)))+(sdB*sdB/float64(len(b))))
	// Compute the p-value.
	p = 2 * (1 - StudentT(float64(len(a)+len(b)-2), t))
	return
}

func Mean(a []float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum / float64(len(a))
}

func StdDev(a []float64) float64 {
	mean := Mean(a)
	var sum float64
	for _, v := range a {
		sum += (v - mean) * (v - mean)
	}
	return math.Sqrt(sum / float64(len(a)))
}

// This function computes the Student's t-distribution.
func StudentT(nu float64, x float64) float64 {
	// Compute the first part of the function.
	a := Gamma((nu + 1) / 2)
	// Compute the second part of the function.
	b := Gamma(nu / 2)
	// Compute the third part of the function.
	c := math.Pow(nu*math.Pi, 0.5)
	// Compute the fourth part of the function.
	d := math.Pow(1+(x*x/nu), -(nu+1)/2)
	// Return the result.
	return a / (b * c * d)
}

// This function computes the Gamma function.
func Gamma(x float64) float64 {
	// Compute the first part of the function.
	a := math.Pow(x, x-0.5)
	// Compute the second part of the function.
	b := math.Exp(-x)
	// Compute the third part of the function.
	c := math.Sqrt(2 * math.Pi)
	// Return the result.
	return a * b * c
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
