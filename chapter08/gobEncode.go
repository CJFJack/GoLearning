package main

import (
	"encoding/gob"
	"os"
	"time"
)

type Task struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

func init() {
	gob.Register(&Task{})
}

func main() {
	start := time.Now()
	end := start.Add(time.Hour * 24)
	tasks := []*Task{
		{
			Id:        1,
			Name:      "task1",
			StartTime: &start,
			EndTime:   &end,
			User:      "cjf",
		},
		{
			Id:        2,
			Name:      "task2",
			StartTime: &start,
			EndTime:   &end,
			User:      "lulu",
		},
	}
	file, _ := os.Create("gob.txt")
	defer file.Close()

	// 定义编码指针对象
	encoder := gob.NewEncoder(file)

	// 编码对象
	encoder.Encode(tasks)
}
