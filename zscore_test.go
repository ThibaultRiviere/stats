package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testZscore(name string, x, avg, stddev, expected float64) {
	Convey(name, func() {
		So(ZScore(x, avg, stddev), ShouldEqual, expected)
	})
}

func testZscoreFromData(name string, datas []float64, x, expected float64) {
	Convey(name, func() {
		So(ZScoreFromData(datas, x), ShouldEqual, expected)
	})
}

func TestZscore(t *testing.T) {
	Convey("Testing zscore", t, func() {
		var ret float64
		var x float64
		var avg float64
		var stddev float64

		ret = 0.
		x = 10.
		avg = 10.
		stddev = 1.
		testZscore("should be equal to 0", x, avg, stddev, ret)

		ret = 5
		x = 15
		avg = 10.
		stddev = 1.
		testZscore("should be equal to 5", x, avg, stddev, ret)

		ret = 10
		x = 20
		avg = 10.
		stddev = 1.
		testZscore("should be equal to 10", x, avg, stddev, ret)
	})
}

func TestZscoreFromData(t *testing.T) {
	Convey("Testing ZScoreFromData", t, func() {
		var ret float64
		var x float64

		datas := []float64{9, 11}

		ret = 0.
		x = 10.
		testZscoreFromData("should be equal to 0", datas, x, ret)

		ret = 5.
		x = 15.
		testZscoreFromData("should be equal to 5", datas, x, ret)

		ret = 10.
		x = 20.
	})
}
