package main

import (
	"fmt"
)

type Sender interface {
}

type Email interface {
}

type Sms interface {
}

type EmailSender struct {
	addr string
}

func (s *EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给：%s，内容：%s\n", to, msg)
	return nil
}

func (s *EmailSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件给：%s，内容：%s\n", to, msg)
	return nil
}

func (s *EmailSender) SendCc(to string, cc string, msg string) error {
	fmt.Printf("发送邮件给：%s，抄送给：%s，内容：%s\n", to, cc, msg)
	return nil
}

type SmsSender struct {
	api string
}

func (s *SmsSender) Send(to string, msg string) error {
	fmt.Printf("发送短信给：%s，内容：%s\n", to, msg)
	return nil
}

func (s *SmsSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送短信给：%s，内容：%s\n", to, msg)
	return nil
}

func (s *SmsSender) SendCc(to string, cc string, msg string) error {
	fmt.Printf("发送短信给：%s，抄送给：%s，内容：%s\n", to, cc, msg)
	return nil
}

func PrintConfig(sender Sender) {
	if obj, ok := sender.(EmailSender); ok {
		fmt.Printf("%#v, %t\n", obj.addr, ok)
	}

	if obj, ok := sender.(*SmsSender); ok {
		fmt.Printf("%#v, %t\n", obj.api, ok)
	}
}

func PrintConfigV2(sender Sender) {
	switch v := sender.(type) {
	case EmailSender:
		fmt.Printf("%#v\n", v.addr)
	case *SmsSender:
		fmt.Printf("%#v\n", v.api)
	default:
		fmt.Println("类型错误")
	}
}

func main() {
	var email Email
	email = EmailSender{"smtp.qq.com"}

	var sms Sms
	sms = &SmsSender{"sms.tengxun.com"}

	PrintConfigV2(email)
	PrintConfigV2(sms)
}
