package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	NPDF_0       = 0.3989422804014327
	NPDF_1       = 0.24197072451914337
	NCDF_0       = 0.5
	NCDF_1       = 0.8413447460685429
	NCDF_MINUS_1 = 0.15865525393145707
)

func TestPDF(t *testing.T) {
	Convey("Testing probability density function", t, func() {
		tests := []dataZscore{
			dataZscore{"with 0", 0, 0, 1, NPDF_0},
			dataZscore{"with 1", 1, 0, 1, NPDF_1},
			dataZscore{"with -1", 1, 0, 1, NPDF_1},
		}
		for _, test := range tests {
			testZscore(test, PDF)
		}
	})

	Convey("Testing normal probability density function", t, func() {
		tests := []dataX{
			dataX{"with 0", 0, NPDF_0},
			dataX{"with 1", 1, NPDF_1},
			dataX{"with -1", -1, NPDF_1},
		}
		for _, test := range tests {
			testX(test, NormalPDF)
		}
	})

	Convey("Testing cumulative distribution function", t, func() {
		tests := []dataZscore{
			dataZscore{"with 0", 0, 0, 1, NCDF_0},
			dataZscore{"with 1", 1, 0, 1, NCDF_1},
			dataZscore{"with -1", -1, 0, 1, NCDF_MINUS_1},
		}
		for _, test := range tests {
			testZscore(test, CDF)
		}
	})

	Convey("Testing normal cumulative distribution function", t, func() {
		tests := []dataX{
			dataX{"with 0", 0, NCDF_0},
			dataX{"with 1", 1, NCDF_1},
			dataX{"with -1", -1, NCDF_MINUS_1},
		}
		for _, test := range tests {
			testX(test, NormalCDF)
		}
	})

	Convey("Testing probability between", t, func() {
		tests := []intervalTest{
			intervalTest{"for 68%", -1, 1, 0, 1, 0.6826894921370859},
			intervalTest{"for 95%", -1.96, 1.96, 0, 1, 0.9500042097035593},
			intervalTest{"for ~100% ", -5, 5, 0, 1, 0.9999994266968564},
			intervalTest{"for ~100% revert", 5, -5, 0, 1, 0.9999994266968564},
		}
		for _, test := range tests {
			testInterval(test, ProbabilityBetween)
		}
	})

	Convey("Testing normalize probability between", t, func() {
		tests := []intervalNormalizeTest{
			intervalNormalizeTest{"for 68%", -1, 1, 0.6826894921370859},
			intervalNormalizeTest{"for 95%", -1.96, 1.96, 0.9500042097035593},
			intervalNormalizeTest{"for ~100% ", -5, 5, 0.9999994266968564},
			intervalNormalizeTest{"for ~100% revert", 5, -5, 0.9999994266968564},
		}
		for _, test := range tests {
			testNormalizeInterval(test, NormalProbabilityBetween)
		}
	})
}
