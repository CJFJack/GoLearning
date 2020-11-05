package main

import (
	"fmt"
	"time"
)

type user struct {
	id   int
	name string
	age  int
}

type task struct {
	name      string
	startTime *time.Time
	endTime   *time.Time
	status    int
	*user
}

func NewUser(id int, name string, age int) *user {
	return &user{
		id:   id,
		name: name,
		age:  age,
	}
}

func NewTask(name string, startTime, endTime *time.Time, status int, user *user) *task {
	return &task{
		name:      name,
		startTime: startTime,
		endTime:   endTime,
		status:    1,
		user:      user,
	}
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) GetName() string {
	return u.name
}

func (t *task) SetName(name string) {
	t.name = name
}

func (t *task) GetName() string {
	return t.name
}

func main() {
	//startTime := time.Now()
	//endTime := startTime.Add(24 * time.Hour)
	//user := NewUser(1, "cjf", 30)
	//task := NewTask("todo homework", &startTime, &endTime, 1, user)
	//
	//fmt.Println(task.GetName())
	//fmt.Println(task.user.GetName())
	//
	//task.SetName("study go")
	//task.user.SetName("lulu")
	//
	//fmt.Println(task.GetName())
	//fmt.Println(task.user.GetName())

	task2 := task{}
	task3 := &task{}

	// 方法值
	//method2 := task2.SetName
	//method3 := task3.SetName
	//
	//method2("abc")
	//method3("123")
	//
	//fmt.Println(task2)
	//fmt.Println(task3)

	// 方法表达式
	method2 := task.SetName    // 如果SetName方法的接收者是指针类型，则不允许
	method3 := (*task).SetName //

	method2(task2, "cjf")
	method3(task3, "lulu")

	fmt.Println(task2)
	fmt.Println(task3)
}
