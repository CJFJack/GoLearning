package main

import "fmt"

type SingleSender interface {
	Send(string, string) error
}

type Sender interface {
	SingleSender
	SendAll([]string, string) error
}

type EmailSender struct {
}

func (e *EmailSender) SendAll(to []string, msg string) error {
	return nil
}

func (e *EmailSender) Send(to, msg string) error {
	return nil
}

func main() {
	var sender Sender = &EmailSender{}
	fmt.Println(sender)
}
