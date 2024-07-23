package components

type Component interface {
	Render()
}

type ListView struct {
	children []Component
}

func (c *ListView) Render() {
	for _, c := range c.children {
		c.Render()
	}
}
