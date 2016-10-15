package coin

import "math/rand"

// Face a string type
type Face string

// const head or tail
const (
	H Face = "head"
	T Face = "tail"
)

// Coin two faces, top or bottom.  Each side can by a head or a tail
type Coin struct {
	Top    Face
	Bottom Face
}

//var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Toss flip coin and return top side
func (c *Coin) Toss() Face {
	r := rand.Intn(2)
	if r == 1 {
		t := c.Top
		c.Top = c.Bottom
		c.Bottom = t
	}
	return c.Top
}
