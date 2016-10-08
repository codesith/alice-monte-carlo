package coin

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCoin(t *testing.T) {
	Convey("When a coin is created", t, func() {
		c := Coin{H, T}
		Convey("Top is head, bottom is tail", func() {
			So(c.Top, ShouldEqual, H)
			So(c.Bottom, ShouldEqual, T)
		})
		Convey("Toss it many times, 50% chance shoud have tail on the top", func() {
			count := 1000000
			tailCount := 0
			for i := 0; i < count; i++ {
				if c.Toss() == T {
					tailCount++
				}
			}
			t := float64(tailCount)
			c := float64(count)
			So(t/c, ShouldAlmostEqual, 0.5, 0.001)
		})
	})
}
