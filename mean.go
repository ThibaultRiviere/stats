package stats

// Mean gets the number expressing the central or typical value in a set of data
func Mean(datas []float64) float64 {
	length := len(datas)
	if length == 0 {
		return 0
	}
	ret := float64(0)
	for i := 0; i < length; i++ {
		ret += datas[i]
	}
	return ret / float64(length)
}
