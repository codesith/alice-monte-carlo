package alice

type Side bool

const (
	H Side = true
	T Side = false
)

type Coin struct {
	top    Side
	bottom Side
}

// flip coin and return top side
func (c *Coin) toss() Side {
	r := random.Intn(2)
	if r == 1 {
		t := c.top
		c.top = c.bottom
		c.bottom = t
	}
	return c.top
}
