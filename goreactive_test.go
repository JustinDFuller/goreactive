package goreactive

import (
	"github.com/justindfuller/goreactive/component"
	"github.com/justindfuller/goreactive/node"
	"github.com/justindfuller/goreactive/tags"

	"testing"
)

const (
	helloWorldExpected = "<p>Hello World</p>"
	childrenExpected   = "<ul><li>one</li><li>two</li></ul>"
)

// helloWorld tests the very basic mechanisms of rendering html.
type helloWorld struct{}

func (_ helloWorld) Render() node.Node {
	return node.New(tags.P, node.Children(node.Text("Hello World")))
}

func TestRenderToString(t *testing.T) {
	result := RenderToString(&helloWorld{})

	if result != helloWorldExpected {
		t.Fatalf("Expected '%s'. Got '%s'.", helloWorldExpected, result)
	}
}

// withChildren tests the mechanisms for rendering child components.
type withChildren struct{}

func (_ withChildren) Render() node.Node {
	return node.New(
		tags.Ul,
		node.Children(
			node.New(
				tags.Li,
				node.Children(node.Text("one")),
			),
			node.New(
				tags.Li,
				node.Children(node.Text("two")),
			),
		),
	)
}

func TestChildrenRenderToString(t *testing.T) {
	result := RenderToString(&withChildren{})

	if result != childrenExpected {
		t.Fatalf("Expected '%s'. Got '%s'.", childrenExpected, result)
	}
}

type childComponent struct {
	text string
}

func (c childComponent) Render() node.Node {
	return node.New(
		tags.Div,
		node.Children(node.Text(c.text)),
	)
}

type withChildComponent struct{}

func (_ withChildComponent) Render() node.Node {
	return component.New(
		childComponent{
			text: "Child component props!",
		},
	)
}
