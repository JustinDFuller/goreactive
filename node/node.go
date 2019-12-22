package node

import (
  "github.com/justindfuller/goreactive/props"

  "fmt"
)

type Node interface {
  ToString() string
}

type Element struct {
  Type string
  Children []Node
}

func (e *Element) ToString() string {
  var channels []chan string
  var children string

  for _, child := range e.Children {
    channel := make(chan string)

    go func(child Node) {
      channel <- child.ToString()
    }(child)

    channels = append(channels, channel)
  }

  for _, channel := range channels {
    str := <-channel
    children += str
  }

  return fmt.Sprintf("<%s>%s</%s>", e.Type, children, e.Type)
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

func (e *TextElement) ToString() string {
  return e.text
}

func Text(text string) Node {
  return &TextElement{
    text: text,
  }
}
