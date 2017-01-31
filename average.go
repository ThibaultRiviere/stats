package stats

// Average gets the number expressing the central or typical value in a set of data
func Average(datas []float32) float32 {
	length := len(datas)
	if length == 0 {
		return 0
	}
	ret := float32(0)
	for i := 0; i < length; i++ {
		ret += datas[i]
	}
	return ret / float32(length)
}
