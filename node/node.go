package node

import (
	"github.com/justindfuller/goreactive/tag"
)

type Node interface {
	ToString() string
}

func New(nodetype tag.Tag, children []Node) Node {
	return &Element{
		Tag:      nodetype,
		Children: children,
	}
}

func toString(channel chan<- string, node Node) {
	channel <- node.ToString()
}

func Children(children ...Node) []Node {
	var childNodes []Node

	for _, child := range children {
		childNodes = append(childNodes, child)
	}

	return childNodes
}
