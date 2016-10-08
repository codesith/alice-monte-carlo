package pocket

import (
	"alice/coin"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomPick(t *testing.T) {
	Convey("When a pocket is created", t, func() {
		p := new(Pocket)
		Convey("It is empty", func() {
			So(p.CountCoins(), ShouldEqual, 0)
		})
		Convey("Add one coin", func() {
			p.Add(coin.Coin{coin.H, coin.T})
			Convey("There should be one coin", func() {
				So(p.CountCoins(), ShouldEqual, 1)
			})
			Convey("Randomly pick one", func() {
				c := p.PickRandom()
				Convey("Should be the same coin", func() {
					So(c.Top, ShouldEqual, "head")
					So(c.Bottom, ShouldEqual, "tail")
				})
				Convey("Pocket is empty", func() {
					So(p.CountCoins(), ShouldEqual, 0)
				})
			})
		})
		Convey("Add two coins, pick one Randomly, and repeat many times, each coin gets 50% chance", func() {
			coins := []coin.Coin{
				coin.Coin{coin.H, coin.H},
				coin.Coin{coin.T, coin.T}}
			count := 1000000
			tailCount := 0
			for i := 0; i < count; i++ {
				p.AddAll(coins)
				c := p.PickRandom()
				if c.Top == coin.T {
					tailCount++
				}
				p.RemoveAll()
			}
			t := float64(tailCount)
			c := float64(count)
			So(t/c, ShouldAlmostEqual, 0.5, 0.001)
		})
	})
}
