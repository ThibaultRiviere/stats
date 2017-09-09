package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testMean(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(Mean(datas), ShouldEqual, expected)
	})
}

func TestMean(t *testing.T) {
	Convey("Testing Mean", t, func() {
		var ret float64
		var datas []float64

		ret = float64(42)
		datas = []float64{32.2, 32.2, 42, 51.8, 51.8}
		testMean("With simple values", datas, ret)

		ret = float64(0)
		datas = []float64{}
		testMean("Without values", datas, ret)
	})
}
