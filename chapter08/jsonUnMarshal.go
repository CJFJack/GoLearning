package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Task struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

func main() {
	tasks := []Task{}
	jsonText, _ := ioutil.ReadFile("task.inputJson")

	json.Unmarshal(jsonText, &tasks)

	for _, task := range tasks {
		fmt.Printf("%#v\n", task)
	}
}
