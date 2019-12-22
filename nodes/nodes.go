package nodes

import (
  "github.com/justindfuller/goreactive/props"

  "fmt"
)

type Node interface {
  TagOpen() string
  TagClose() string
  RenderChildren() string
}

type Element struct {
  Type string
  Children []Node
}

func (e *Element) TagOpen() string {
  return fmt.Sprintf("<%s>", e.Type)
}

func (e *Element) TagClose() string {
  return fmt.Sprintf("</%s>", e.Type)
}

func (e *Element) RenderChildren() string {
  children := ""

  for _, child := range e.Children {
    children += fmt.Sprintf("%s%s%s", child.TagOpen(), child.RenderChildren(), child.TagClose())
  }

  return children
}

func Create(nodetype string, props *props.Props, children []Node) Node {
  return &Element{
    Type: nodetype,
    Children: children,
  }
}

func Children(children ...Node) []Node {
  var childNodes []Node

  for _, child := range children {
    childNodes = append(childNodes, child)
  }

  return childNodes
}

type TextElement struct {
  text string 
} 

func (_ *TextElement) TagOpen() string {
  return ""
}

func (_ *TextElement) TagClose() string {
  return ""
}

func (e *TextElement) RenderChildren() string {
  return e.text
}

func Text(text string) Node {
  return &TextElement{
    text: text,
  }
}
