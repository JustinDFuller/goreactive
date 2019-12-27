package goreactive

import (
	"testing"
)

func TestParseElement(t *testing.T) {
	str := "<div>Child text</div>"
	expected := "node.New(tag.Div, node.Children(node.Text(\"Child Text\")))"

	if output := Parse(str); output != expected {
		t.Fatalf("Expected %s. Got %s.", expected, output)
	}
}
