package main

import (
	"bytes"
	"encoding/json"
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

func main() {
	start := time.Now()
	end := time.Now().Add(time.Hour * 24)
	tasks := []Task{
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

	ctx, err := json.Marshal(tasks)
	fmt.Println(json.Valid(ctx))
	fmt.Println(string(ctx), err)

	buffer := bytes.Buffer{}
	json.Indent(&buffer, ctx, "", "    ")

	file, _ := os.Create("task.inputJson")
	buffer.WriteTo(file)

	jsonText, _ := json.MarshalIndent(tasks, "", "    ")
	fmt.Println(string(jsonText))

}
