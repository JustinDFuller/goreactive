package goreactive

import (
	"github.com/justindfuller/goreactive/component"
	"github.com/justindfuller/goreactive/node"
	"github.com/justindfuller/goreactive/tag"

	"testing"
)

const (
	ExpectedHelloWorld     = "<p>Hello World</p>"
	ExpectedChildren       = "<ul><li>one</li><li>two</li></ul>"
	ExpectedChildComponent = "<div>Child component props!</div>"
)

// helloWorld tests the very basic mechanisms of rendering html.
type helloWorld struct{}

func (_ *helloWorld) Render() node.Node {
	return node.New(tag.P, node.Children(node.Text("Hello World")))
}

func TestRenderToString(t *testing.T) {
	result := RenderToString(&helloWorld{})

	if result != ExpectedHelloWorld {
		t.Fatalf("Expected '%s'. Got '%s'.", ExpectedHelloWorld, result)
	}
}

// withChildren tests the mechanisms for rendering child components.
type withChildren struct{}

func (_ *withChildren) Render() node.Node {
	return node.New(
		tag.Ul,
		node.Children(
			node.New(
				tag.Li,
				node.Children(node.Text("one")),
			),
			node.New(
				tag.Li,
				node.Children(node.Text("two")),
			),
		),
	)
}

func TestChildrenRenderToString(t *testing.T) {
	result := RenderToString(&withChildren{})

	if result != ExpectedChildren {
		t.Fatalf("Expected '%s'. Got '%s'.", ExpectedChildren, result)
	}
}

type childComponent struct {
	text string
}

func (c childComponent) Render() node.Node {
	return node.New(
		tag.Div,
		node.Children(node.Text(c.text)),
	)
}

type withChildComponent struct{}

func (_ *withChildComponent) Render() node.Node {
	return component.New(
		childComponent{
			text: "Child component props!",
		},
	)
}

func TestChildComponent(t *testing.T) {
	result := RenderToString(&withChildComponent{})

	if result != ExpectedChildComponent {
		t.Fatalf("Expected '%s'. Got '%s'.", ExpectedChildComponent, result)
	}
}
