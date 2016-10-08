package main

import "alice/coin"
import "alice/pocket"
import "fmt"

func main() {
	p := new(pocket.Pocket)
	coins := []coin.Coin{
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.H, coin.T},
		coin.Coin{coin.T, coin.T},
		coin.Coin{coin.T, coin.T}}
	p.Add(coins)
	c := p.PickRandom()
	fmt.Println(c)
	fmt.Println(p.CountCoins())
	fmt.Println(c.Toss())
}
