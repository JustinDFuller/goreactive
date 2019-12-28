package node

import (
	"github.com/justindfuller/goreactive/tag"

	"bytes"
	"fmt"
	"io"
)

type (
  Element struct {
    Tag      tag.Tag
    Children []Node
  }

  TextElement struct {
    text string
  }
)

func (e *Element) ToString(out io.Writer) {
	var channels []chan io.ReadWriter

	fmt.Fprintf(out, "<%s>", e.Tag)

	for _, child := range e.Children {
		var childString bytes.Buffer
		channel := make(chan io.ReadWriter)
		channels = append(channels, channel)
		go toString(channel, child, &childString)
	}

	for _, channel := range channels {
		child := <-channel
		io.Copy(out, child)
	}

	fmt.Fprintf(out, "</%s>", e.Tag)
}

func (e *TextElement) ToString(out io.Writer) {
	fmt.Fprint(out, e.text)
}

func Text(text string) Node {
	return &TextElement{
		text: text,
	}
}
