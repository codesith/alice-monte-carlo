package alice

import "time"

import "math/rand"

type Side bool

const (
	H Side = true
	T Side = false
)

type Coin struct {
	top    Side
	bottom Side
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

// flip coin and return top side
func (c *Coin) toss() Side {
	r := r1.Intn(2)
	if r == 1 {
		t := c.top
		c.top = c.bottom
		c.bottom = t
	}
	return c.top
}
