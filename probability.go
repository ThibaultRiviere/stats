package stats

import "math"

// PDF normalize x using the mean and the standard deviation and return the PDF
// https://en.wikipedia.org/wiki/Probability_density_function
func PDF(x, avg, stddev float64) float64 {
	return math.Exp(-math.Pow(((x-avg)/stddev), 2)/2) / (stddev * math.Sqrt(2*math.Pi))
}

// return the PDF with z already normalized
// https://en.wikipedia.org/wiki/Probability_density_function
func NormalPDF(z float64) float64 {
	return math.Exp(-math.Pow(z, 2)/2) / math.Sqrt(2*math.Pi)
}

// CDF return the CDF using the mean and the standard deviation given
// https://en.wikipedia.org/wiki/Cumulative_distribution_function#Definition
func CDF(x, avg, stddev float64) float64 {
	return (1 + math.Erf(((x-avg)/stddev)/math.Sqrt2)) / 2
}

// NormalCDF return the CDF with z as already normalize
// https://en.wikipedia.org/wiki/Cumulative_distribution_function#Definition
func NormalCDF(z float64) float64 {
	return (1 + math.Erf(z/(math.Sqrt2))) / 2
}

// ProbabilityBetween return after normalize the probability
// that X will take a value between x and y.
func ProbabilityBetween(x, y, avg, stddev float64) float64 {
	return math.Abs(CDF(x, avg, stddev) - CDF(y, avg, stddev))
}

// ProbabilityBetween return the probability that X will take a value
// between x and y with z1 and z2 already normalized.
func NormalProbabilityBetween(z1, z2 float64) float64 {
	return math.Abs(NormalCDF(z1) - NormalCDF(z2))
}
