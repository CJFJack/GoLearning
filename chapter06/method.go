package main

import (
	"fmt"
	"time"
)

type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      string
}

func (task *Task) SetName(user string) {
	task.user = user
}

func main() {
	task := &Task{
		name: "toDoList",
	}
	task.SetName("cjf")
	fmt.Println(task)
	(*task).SetName("lulu")
	fmt.Println(task)

	var point *int
	fmt.Println(*point)
}
