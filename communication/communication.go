package communication

type Communication struct {
	C Chan
}

func (c *Communication) Request(i interface{}) {
	c.C.In(i)
}

func (c *Communication) Resp() interface{} {
	return c.C.Out()
}
