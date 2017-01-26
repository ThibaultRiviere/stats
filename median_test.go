package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func testMedian(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(Median(datas), ShouldEqual, expected)
	})
}

func TestMedian(t *testing.T) {
	Convey("Testing Median", t, func() {
		var ret float64
		var datas []float64

		ret = (float64(1.3) + float64(1.35)) / 2

		datas = []float64{1.1, 1.3, 1.35, 1.4}
		testMedian("With even array length and sorted values", datas, ret)

		datas = []float64{1.3, 1.4, 1.35, 1.1}
		testMedian("With even array length and not sorted values", datas, ret)

		ret = float64(1.3)

		datas = []float64{1.1, 1.3, 1.4}
		testMedian("With not even array length and sorted values", datas, ret)

		datas = []float64{1.3, 1.4, 1.1}
		testMedian("With not even array length and not sorted values", datas, ret)

		median := Median([]float64{})
		Convey("With no datas", func() {
			So(math.IsNaN(median), ShouldEqual, true)
		})
	})
}
