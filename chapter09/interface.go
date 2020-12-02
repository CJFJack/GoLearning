package main

import (
	"fmt"
)

// 定义接口
type Sender interface {
	// 定义接口的行为（方法名、参数列表、返回列表）
	Send(string, string) error
	SendAll([]string, string) error
	SendCc(string, string, string) error
}

// 定义结构体
type EmailSender struct {
	addr     string
	port     int
	user     string
	password string
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

func main() {
	// 声明接口变量
	var sender Sender
	fmt.Printf("%T, %#v\n", sender, sender)

	// 给接口变量赋值
	sender = &EmailSender{"smtp.qq.com", 465, "user", "password"}

	// 调用接口方法
	sender.Send("cjf", "hello")
	sender.SendAll([]string{"cjf", "lulu"}, "hello")
	sender.SendCc("cjf", "lulu", "heihei")

}
