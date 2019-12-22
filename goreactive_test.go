package goreactive

import (
  "github.com/justindfuller/goreactive/tags"
  "github.com/justindfuller/goreactive/nodes"
  "github.com/justindfuller/goreactive/props"

  "testing"
)

const (
  helloWorldExpected = "<p>Hello World</p>"
  childrenExpected = "<ul><li>one</li><li>two</li></ul>"
)

// helloWorld tests the very basic mechanisms of rendering html.
type helloWorld struct {}

func (_ helloWorld) Render() nodes.Node {
  return nodes.Create(tags.P, props.New(), nodes.Children(nodes.Text("Hello World")))
}

func TestRenderToString(t *testing.T) {
  result := RenderToString(&helloWorld{})

  if result != helloWorldExpected {
    t.Fatalf("Expected '%s'. Got '%s'.", helloWorldExpected, result)
  }
}

// withChildren tests the mechanisms for rendering child components.
type withChildren struct {}

func (_ withChildren) Render() nodes.Node {
  return nodes.Create(
    tags.Ul,
    props.New(),
    nodes.Children(
      nodes.Create(
        tags.Li,
        props.New(),
        nodes.Children(nodes.Text("one")),
      ),
      nodes.Create(
        tags.Li,
        props.New(),
        nodes.Children(nodes.Text("two")),
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
