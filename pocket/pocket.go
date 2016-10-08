package pocket

import (
	"alice/coin"
	"container/list"
	"math/rand"
	"time"
)

// Pocket a container or coins
type Pocket struct {
	coins list.List
}

// Add coins into a pocket
func (p *Pocket) Add(coins []coin.Coin) {
	for _, c := range coins {
		p.coins.PushBack(c)
	}
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// PickRandom randomly pick up a coin
func (p *Pocket) PickRandom() (c coin.Coin) {
	index := random.Intn(p.coins.Len())
	i := 0
	for e := p.coins.Front(); e != nil; e = e.Next() {
		if i == index {
			c := p.coins.Remove(e)
			return c.(coin.Coin)
		}
		i++
	}
	return
}

// CountCoins counts how many coins are in a pocket
func (p Pocket) CountCoins() (count int) {
	return p.coins.Len()
}
