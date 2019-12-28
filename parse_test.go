package goreactive

import (
	"strings"
	"testing"
)

func TestParseElement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple div",
			input:    "<div>Child text</div>",
			expected: "node.New(tag.Div, node.Children(node.Text(\"Child Text\")))",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result strings.Builder
			Parse(test.input, &result)

			if output := result.String(); output != test.expected {
				t.Fatalf("Expected %s. Got %s.", test.expected, output)
			}
		})
	}
}
