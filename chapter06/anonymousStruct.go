package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
	tel  string
}

type Task struct {
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	*User
}

func main() {
	task := Task{
		name:   "todolist",
		status: 1,
		User: &User{
			"cjf",
			30,
			"1351021479",
		},
	}

	task2 := task

	task.User.age = 29
	fmt.Printf("%#v\n", task.User)

	fmt.Printf("%#v\n", task2.User)

}
