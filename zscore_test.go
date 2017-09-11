package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestZscore(t *testing.T) {
	Convey("Testing zscore", t, func() {
		tests := []dataZscore{
			dataZscore{"should be equal to 0 ", 10, 10, 1, 0},
			dataZscore{"should be equal to 5 ", 15, 10, 1, 5},
			dataZscore{"should be equal to 10", 20, 10, 1, 10},
		}
		for _, test := range tests {
			testZscore(test, ZScore)
		}
	})
}
