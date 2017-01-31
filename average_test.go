package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testAverage(name string, datas []float32, expected float32) {
	Convey(name, func() {
		So(Average(datas), ShouldEqual, expected)
	})
}

func TestAverage(t *testing.T) {
	Convey("Testing Average", t, func() {
		var ret float32
		var datas []float32

		ret = float32(42)
		datas = []float32{32.2, 32.2, 42, 51.8, 51.8}
		testAverage("With simple values", datas, ret)

		ret = float32(0)
		datas = []float32{}
		testAverage("Without values", datas, ret)
	})
}
