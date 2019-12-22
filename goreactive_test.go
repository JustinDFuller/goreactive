package goreactive

import (
  "testing"
)

const (
  helloWorldExpected = "<p>Hello World</p>"
)

type helloWorld struct {}

func (_ helloWorld) Render() *Node {
  return CreateElement("p", &Props{}, "Hello World")
}

func TestRenderToString(t *testing.T) {
  result := RenderToString(&helloWorld{})

  if result != helloWorldExpected {
    t.Fatalf("Expected '%s'. Got '%s'.", helloWorldExpected, result)
  }
}
