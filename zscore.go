package stats

// Zscore is the distance from the mean in stddev
func ZScore(x float64, avg float64, stddev float64) float64 {
	return (x - avg) / stddev
}

// ZScoreFromData will compute the mean and the stddev before computing
// the zscore
func ZScoreFromData(datas []float64, x float64) float64 {
	return ZScore(x, Mean(datas), StdDev(datas))
}
