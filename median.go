package stats

import (
	"math"
	"sort"
)

// Median gets a simple measure of central tendency from the slices of float64
func Median(datas []float64) float64 {
	length := len(datas)
	if length == 0 {
		return math.NaN()
	}
	sort.Float64s(datas)
	if length%2 != 0 {
		return datas[length/2]
	}
	return (datas[length/2-1] + datas[length/2]) / 2
}
