package node

import (
  "github.com/justindfuller/goreactive/tags"

  "fmt"
  "strings"  
)

type Element struct {
	Tag      tag.Tag
	Children []Node
}

func (e *Element) ToString() string {
	var channels []chan string
	var children strings.Builder

	for _, child := range e.Children {
		channel := make(chan string)
		channels = append(channels, channel)
		go toString(channel, child)
	}

	for _, channel := range channels {
		children.WriteString(<-channel)
	}

	return fmt.Sprintf("<%s>%s</%s>", e.Tag, children.String(), e.Tag)
}

type TextElement struct {
	text string
}

func (e *TextElement) ToString() string {
	return e.text
}

func Text(text string) Node {
	return &TextElement{
		text: text,
	}
}
