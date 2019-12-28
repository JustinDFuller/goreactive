package goreactive

import (
	"github.com/justindfuller/goreactive/component"
	"github.com/justindfuller/goreactive/node"
	"github.com/justindfuller/goreactive/tag"

	"strings"
	"testing"
)

type RenderToStringTest struct {
	name     string
	expected string
	input    component.Component
}

func TestRenderToString(t *testing.T) {
	tests := []RenderToStringTest{
		{
			name:     "Hello World",
			expected: ExpectedHelloWorld,
			input:    &helloWorld{},
		},
		{
			name:     "With Children",
			expected: ExpectedChildren,
			input:    &withChildren{},
		},
		{
			name:     "Child Component",
			expected: ExpectedChildComponent,
			input:    &withChildComponent{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result strings.Builder
			RenderToString(test.input, &result)

			if result := result.String(); result != test.expected {
				t.Fatalf("Expected '%s'. Got '%s'.", test.expected, result)
			}
		})
	}
}

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
