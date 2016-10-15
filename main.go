package main

import (
	"alice/coin"
	"alice/pocket"
	"fmt"
	"time"
)

func main() {
	coins := []coin.Coin{
		coin.Coin{coin.H, coin.T},
		coin.Coin{coin.H, coin.T},
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.H, coin.H},
		coin.Coin{coin.T, coin.T}}

	trials := 1000000000
	routineCount := 10
	overall := 0
	success := 0
	var o chan int = make(chan int, routineCount)
	var s chan int = make(chan int, routineCount)
	subtrials := trials / routineCount
	start := time.Now()
	for i := 0; i < routineCount; i++ {
		go trial(coins, subtrials, o, s)
		fmt.Printf("Start routine %d\n", i)
	}
	for i := 0; i < routineCount; i++ {
		overall = overall + <-o
		success = success + <-s
	}
	elapsed := time.Since(start)
	fmt.Printf("\rFinal result is %.2f%%\n", float64(success)/float64(overall)*100)
	fmt.Printf("Processing time is %s\n", elapsed)
}

func trial(coins []coin.Coin, trials int, overall chan int, success chan int) {
	o := 0
	s := 0
	for i := 0; i < trials; i++ {
		p := new(pocket.Pocket)
		p.AddAll(coins)
		// draw a single coin
		c := p.PickRandom()
		// toss it
		if c.Toss() == coin.H {
			o++
			// draw another coin
			c = p.PickRandom()
			if c.Toss() == coin.H {
				s++
			}
		}
	}
	overall <- o
	success <- s
}
