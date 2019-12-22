package component

import (
  "github.com/justindfuller/goreactive/nodes"
)

type Component interface {
  Render() nodes.Node
}
