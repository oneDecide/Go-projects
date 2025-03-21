package main

import "fmt"

type Message struct {
	text string
}

func NewMessage(newText string) Message {
	m := Message{text: newText}
	return m
}

func (m Message) PrintMessage() {
	fmt.Println(m.text)
}
