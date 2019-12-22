package component

import (
  "github.com/justindfuller/goreactive/node"
)

type Component interface {
  Render() node.Node
}
