package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMedian(t *testing.T) {
	Convey("Testing Median", t, func() {

		tests := []dataTest{
			dataTest{
				"With even array length and sorted values",
				[]float64{1.1, 1.3, 1.35, 1.4},
				(float64(1.3) + float64(1.35)) / 2,
			},
			dataTest{
				"With even array length and not sorted values",
				[]float64{1.3, 1.4, 1.35, 1.1},
				(float64(1.3) + float64(1.35)) / 2,
			},
			dataTest{
				"With not even array length and sorted values",
				[]float64{1.1, 1.3, 1.4},
				1.3,
			},
			dataTest{
				"With not even array length and not sorted values",
				[]float64{1.3, 1.4, 1.1},
				1.3,
			},
		}
		for _, test := range tests {
			testData(test, Median)
		}

	})
}
