package goreactive

import (
	"testing"
)

func BenchmarkRenderToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderToString(&helloWorld{})
	}
}

func BenchmarkChildrenRenderToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderToString(&withChildren{})
	}
}
