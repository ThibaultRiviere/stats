package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testVariance(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(Variance(datas), ShouldEqual, expected)
	})
}

func testBesselVariance(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(BesselVariance(datas), ShouldEqual, expected)
	})
}

func testStdDev(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(StdDev(datas), ShouldEqual, expected)
	})
}

func testBesselStdDev(name string, datas []float64, expected float64) {
	Convey(name, func() {
		So(BesselStdDev(datas), ShouldEqual, expected)
	})
}

func testSE(name string, stddev float64, n int, expected float64) {
	Convey(name, func() {
		So(SE(stddev, n), ShouldEqual, expected)
	})
}

func TestVariance(t *testing.T) {
	Convey("Testing Bessel Variance", t, func() {
		var ret float64
		var datas []float64

		ret = float64(0)
		datas = []float64{}
		testBesselVariance("Without values", datas, ret)

		ret = float64(1)
		datas = []float64{3, 4, 5}
		testBesselVariance("with variance of 1", datas, ret)

		ret = float64(0)
		datas = []float64{1, 1, 1, 1}
		testBesselVariance("with no variance", datas, ret)
	})
}

func TestBesselVariance(t *testing.T) {
	Convey("Testing Variance", t, func() {
		var ret float64
		var datas []float64

		ret = float64(0)
		datas = []float64{}
		testBesselVariance("Without values", datas, ret)

		ret = float64(1)
		datas = []float64{3, 4, 5}
		testBesselVariance("with variance of 1", datas, ret)

		ret = float64(0)
		datas = []float64{1, 1, 1, 1}
		testBesselVariance("with no variance", datas, ret)
	})
}

func TestStdDev(t *testing.T) {
	Convey("Testing Stddev", t, func() {
		var ret float64
		var datas []float64

		ret = float64(0)
		datas = []float64{}
		testStdDev("without value", datas, ret)

		ret = float64(1)
		datas = []float64{3, 5}
		testStdDev("with stddev of 1", datas, ret)

		ret = float64(0)
		datas = []float64{1, 1, 1, 1}
		testStdDev("with no stddev", datas, ret)
	})
}

func TestBesselStdDev(t *testing.T) {
	Convey("Testing bessel Stddev", t, func() {
		var ret float64
		var datas []float64

		ret = float64(0)
		datas = []float64{}
		testBesselStdDev("without value", datas, ret)

		ret = float64(1)
		datas = []float64{3, 4, 5}
		testBesselStdDev("with stddev of 1", datas, ret)

		ret = float64(0)
		datas = []float64{1, 1, 1, 1}
		testBesselStdDev("with no stddev", datas, ret)
	})
}

func TestSE(t *testing.T) {
	Convey("Testing standard Error", t, func() {
		var ret float64
		var stddev float64
		var n int

		ret = 1.
		stddev = 1.
		n = 1
		testSE("should be equal to 1", stddev, n, ret)

		ret = 10
		stddev = 100
		n = 100
		testSE("should be equal to 10", stddev, n, ret)
	})
}
