package pocket

import (
	"alice/coin"
	"container/list"
	"math/rand"
)

// Pocket a container or coins
type Pocket struct {
	coins list.List
}

// AddAll coins into a pocket
func (p *Pocket) AddAll(coins []coin.Coin) {
	for _, c := range coins {
		p.coins.PushBack(c)
	}
}

// Add a coin into a pocket
func (p *Pocket) Add(c coin.Coin) {
	p.coins.PushBack(c)
}

// Remove all coin
func (p *Pocket) RemoveAll() {
	p.coins.Init()
}

//var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// PickRandom randomly pick up a coin
func (p *Pocket) PickRandom() (c coin.Coin) {
	index := rand.Intn(p.coins.Len())
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
