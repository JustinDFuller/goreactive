package node

import (
	"github.com/justindfuller/goreactive/tag"

	"io"
)

type Node interface {
	ToString(out io.Writer)
}

func New(nodetype tag.Tag, children []Node) Node {
	return &Element{
		Tag:      nodetype,
		Children: children,
	}
}

func toString(done chan<- io.ReadWriter, node Node, out io.ReadWriter) {
	node.ToString(out)
	done <- out
}

func Children(children ...Node) []Node {
	var childNodes []Node

	for _, child := range children {
		childNodes = append(childNodes, child)
	}

	return childNodes
}
