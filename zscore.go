package stats

// Zscore is the distance from the mean in stddev
func ZScore(x, avg, stddev float64) float64 {
	return (x - avg) / stddev
}
