package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestVariance(t *testing.T) {
	Convey("Testing Bessel Variance", t, func() {
		tests := []dataTest{
			dataTest{"without value", []float64{}, 0},
			dataTest{"with variance of 1", []float64{3, 4, 5}, 1},
			dataTest{"with no variance", []float64{1, 1, 1, 1}, 0},
		}
		for _, test := range tests {
			testData(test, BesselVariance)
		}
	})
}

func TestBesselVariance(t *testing.T) {
	Convey("Testing Variance", t, func() {
		tests := []dataTest{
			dataTest{"without value", []float64{}, 0},
			dataTest{"with variance of 1", []float64{3, 5}, 1},
			dataTest{"with no variance", []float64{1, 1, 1, 1}, 0},
		}
		for _, test := range tests {
			testData(test, Variance)
		}
	})
}

func TestStdDev(t *testing.T) {
	Convey("Testing Stddev", t, func() {
		tests := []dataTest{
			dataTest{"without value", []float64{}, 0},
			dataTest{"with stddev of 1", []float64{3, 5}, 1},
			dataTest{"with no stddev", []float64{1, 1, 1, 1}, 0},
		}
		for _, test := range tests {
			testData(test, StdDev)
		}
	})
}

func TestBesselStdDev(t *testing.T) {
	Convey("Testing bessel Stddev", t, func() {
		tests := []dataTest{
			dataTest{"without value", []float64{}, 0},
			dataTest{"with stddev of 1", []float64{3, 4, 5}, 1},
			dataTest{"with no stddev", []float64{1, 1, 1, 1}, 0},
		}
		for _, test := range tests {
			testData(test, BesselStdDev)
		}
	})
}

func TestSE(t *testing.T) {
	Convey("Testing standard Error", t, func() {
		testSE("should be equal to 1", 1, 1, 1)
		testSE("should be equal to 10", 100, 100, 10)
	})
}
