package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMean(t *testing.T) {
	Convey("Testing Mean", t, func() {
		tests := []dataTest{
			dataTest{"With simple values", []float64{32.2, 32.2, 42, 51.8, 51.8}, 42},
			dataTest{"Without values", []float64{}, 0},
		}
		for _, test := range tests {
			testData(test, Mean)
		}
	})
}
