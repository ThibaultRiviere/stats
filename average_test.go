package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testAverage(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(Average(datas), ShouldEqual, expected)
	})
}

func TestAverage(t *testing.T) {
	Convey("Testing Average", t, func() {
		var ret float64
		var datas []float64

		ret = float64(42)
		datas = []float64{32.2, 32.2, 42, 51.8, 51.8}
		testAverage("With simple values", datas, ret)

		ret = float64(0)
		datas = []float64{}
		testAverage("Without values", datas, ret)
	})
}
