package stats

import "math"

// Variance is the mean of the sum of all square deviation
func Variance(datas []float64) float64 {
	avg := Mean(datas)
	var ret float64 = 0
	for i := 0; i < len(datas); i++ {
		dist := datas[i] - avg
		if dist != 0 {
			ret += math.Pow(dist, 2)
		}
	}
	if ret == 0 {
		return 0.0
	}
	return ret / float64(len(datas))
}

// besselVariance generalize the variance for the population from a simple
// of the population
func BesselVariance(datas []float64) float64 {
	avg := Mean(datas)
	var ret float64 = 0
	for i := 0; i < len(datas); i++ {
		ret += math.Pow(datas[i]-avg, 2)
	}
	return ret / float64(len(datas)-1)
}

// StdDev is the standard deviation, the square root of the variance
func StdDev(datas []float64) float64 {
	return math.Sqrt(Variance(datas))
}

// Bessel Standard deviation generalize the standard deviation for the population
// from a simple of the population
func BesselStdDev(datas []float64) float64 {
	return math.Sqrt(BesselVariance(datas))
}

// SE (standard error) is the standard deviation of its sampling distribution
func SE(stddev float64, n int) float64 {
	return stddev / math.Sqrt(float64(n))
}
