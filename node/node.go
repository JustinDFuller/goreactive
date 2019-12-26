package node

import (
	"github.com/justindfuller/goreactive/tag"

	"fmt"
	"strings"
)

type Node interface {
	ToString() string
}

type Element struct {
	Tag      tag.Tag
	Children []Node
}

func (e *Element) ToString() string {
	var channels []chan string
	var children strings.Builder

	for _, child := range e.Children {
		channel := make(chan string)

		go func(child Node) {
			channel <- child.ToString()
		}(child)

		channels = append(channels, channel)
	}

	for _, channel := range channels {
		children.WriteString(<-channel)
	}

	return fmt.Sprintf("<%s>%s</%s>", e.Tag, children.String(), e.Tag)
}

func New(nodetype tag.Tag, children []Node) Node {
	return &Element{
		Tag:      nodetype,
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
