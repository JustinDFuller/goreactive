package goreactive

import (
	"github.com/justindfuller/goreactive/node"
	"github.com/justindfuller/goreactive/props"
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
	return node.Create(tags.P, props.New(), node.Children(node.Text("Hello World")))
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
	return node.Create(
		tags.Ul,
		props.New(),
		node.Children(
			node.Create(
				tags.Li,
				props.New(),
				node.Children(node.Text("one")),
			),
			node.Create(
				tags.Li,
				props.New(),
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
