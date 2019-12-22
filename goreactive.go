package goreactive

import (
  "github.com/justindfuller/goreactive/component"

  "fmt"
)

func RenderToString(c component.Component) string {
  node := c.Render()
  return fmt.Sprintf("%s%s%s", node.TagOpen(), node.RenderChildren(), node.TagClose())
}

