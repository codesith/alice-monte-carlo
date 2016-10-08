package coin

import (
	"math/rand"
	"time"
)

// Side a string type
type Side string

// const head or tail
const (
	H Side = "head"
	T Side = "tail"
)

// Coin two sides, top or bottom.  Each side can by a head or a tail
type Coin struct {
	Top    Side
	Bottom Side
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Toss flip coin and return top side
func (c *Coin) Toss() Side {
	r := random.Intn(2)
	if r == 1 {
		t := c.Top
		c.Top = c.Bottom
		c.Bottom = t
	}
	return c.Top
}
