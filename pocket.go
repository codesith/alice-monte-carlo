package alice

import "container/list"

type Pocket struct {
	coins list.List
}

// dump coins into a pocket
func (p *Pocket) Dump(coins []Coin) {
	for _, c := range coins {
		p.coins.PushBack(c)
	}
}

// randomly pick up a coin
func (p *Pocket) PickRandom() (c Coin) {
	index := random.Intn(p.coins.Len())
	i := 0
	for e := p.coins.Front(); e != nil; e = e.Next() {
		i++
		if i == index {
			c := p.coins.Remove(e)
			return c.(Coin)
		}
	}
	return
}
