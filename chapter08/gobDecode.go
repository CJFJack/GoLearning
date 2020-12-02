package main

import (
	"encoding/gob"
	"fmt"
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
	var tasks []*Task
	file, _ := os.Open("gob.txt")
	defer file.Close()

	// 定义解码指针对象
	decoder := gob.NewDecoder(file)

	// 解码
	decoder.Decode(&tasks)

	for _, task := range tasks {
		fmt.Printf("%v\n", task)
	}

}
