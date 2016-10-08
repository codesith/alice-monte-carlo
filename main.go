package main

import (
	"alice/coin"
	"alice/pocket"
	"fmt"
)

func main() {
	p := new(pocket.Pocket)
	coins := []coin.Coin{
		coin.Coin{coin.H, coin.T},
		coin.Coin{coin.H, coin.T},
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.T, coin.T}}

	trails := 1000000000
	overall := 0
	success := 0
	for i := 0; i < trails; i++ {
		p.AddAll(coins)
		// draw a single coin
		c := p.PickRandom()
		// toss it
		if c.Toss() == coin.H {
			overall++
			// draw another coin
			c = p.PickRandom()
			if c.Toss() == coin.H {
				success++
			}
		}
		p.RemoveAll()
	}
	fmt.Println(float64(success) / float64(overall))
}
