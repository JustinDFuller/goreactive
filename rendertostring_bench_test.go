package goreactive

import (
	"strings"
	"testing"
)

func BenchmarkRenderToString(b *testing.B) {
	var results strings.Builder

	for i := 0; i < b.N; i++ {
		RenderToString(&helloWorld{}, &results)
	}
}

func BenchmarkChildrenRenderToString(b *testing.B) {
	var results strings.Builder

	for i := 0; i < b.N; i++ {
		RenderToString(&withChildren{}, &results)
	}
}
