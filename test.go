package stats

import . "github.com/smartystreets/goconvey/convey"

//
// Helper for factorize tests
//

type dataTest struct {
	name     string
	datas    []float64
	expected float64
}

type dataZscore struct {
	name                     string
	x, avg, stddev, expected float64
}

type dataX struct {
	name        string
	x, expected float64
}

type f func(data []float64) float64
type x func(x float64) float64
type fZscore func(x, avg, stddev float64) float64

func testData(t dataTest, function f) {
	Convey(t.name, func() {
		So(function(t.datas), ShouldEqual, t.expected)
	})
}

func testX(t dataX, function x) {
	Convey(t.name, func() {
		So(function(t.x), ShouldEqual, t.expected)
	})
}

func testZscore(t dataZscore, function fZscore) {
	Convey(t.name, func() {
		So(function(t.x, t.avg, t.stddev), ShouldEqual, t.expected)
	})
}

func testSE(name string, stddev float64, n int, expected float64) {
	Convey(name, func() {
		So(SE(stddev, n), ShouldEqual, expected)
	})
}

type intervalNormalizeTest struct {
	name           string
	x, y, expected float64
}

type FNI func(x, y float64) float64

type intervalTest struct {
	name                        string
	x, y, avg, stddev, expected float64
}

type FI func(x, y, avg, stddev float64) float64

func testNormalizeInterval(t intervalNormalizeTest, function FNI) {
	Convey(t.name, func() {
		So(function(t.x, t.y), ShouldEqual, t.expected)
	})
}

func testInterval(t intervalTest, function FI) {
	Convey(t.name, func() {
		So(function(t.x, t.y, t.avg, t.stddev), ShouldEqual, t.expected)
	})
}
