package goreactive

import (
	"github.com/justindfuller/goreactive/component"

	"io"
)

func RenderToString(c component.Component, out io.Writer) {
	node := c.Render()
	node.ToString(out)
}

func Parse(str string) string {
	return ""
}
