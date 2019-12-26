package component

import (
	"github.com/justindfuller/goreactive/node"
)

type Component interface {
	Render() node.Node
}

func New(c Component) node.Node {
	return c.Render()
}
