package types

type Counter struct {
	num int
}

func (c *Counter) Init() Counter {
	return Counter{0}
}

func (c *Counter) New() int {
	c.num++
	return c.num - 1
}

func (c Counter) Current() int {
	return c.num
}
func (c Counter) Last() int {
	return c.num - 1
}
