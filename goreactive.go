package goreactive

import (
	"github.com/justindfuller/goreactive/component"

	"fmt"
	"io"
)

func RenderToString(c component.Component, out io.Writer) {
	node := c.Render()
	node.ToString(out)
}

func Parse(str string, out io.Writer) {
	fmt.Fprint(out, "")
}
