package goreactive

import (
	"github.com/justindfuller/goreactive/component"
)

func RenderToString(c component.Component) string {
	node := c.Render()
	return node.ToString()
}
