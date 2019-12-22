package goreactive

import (
  "fmt"
)

type Node struct {
  Type string
  Children []interface{}
}

type Props struct {}

type Component interface {
  Render() *Node
}

func RenderToString(component Component) string {
  node := component.Render()
  return fmt.Sprintf("<%s>%s<%s>", node.Type, node.Children[0], node.Type)
}

func CreateElement(elementType string, props *Props, children string) *Node {
  return &Node{
    Type: elementType,
    Children: []interface{}{children},
  }
}
