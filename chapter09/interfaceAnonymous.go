package main

import "fmt"

type EmailSender struct {
}

func (e EmailSender) Send(to, msg string) error {
	return nil
}

func main() {
	var sender interface {
		Send(string, string) error
	}
	sender = EmailSender{}
	fmt.Println(sender)
}
