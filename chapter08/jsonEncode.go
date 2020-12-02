package main

import (
	"bufio"
	"encoding/json"
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

	file, _ := os.Create("task2.json")
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	encoder := json.NewEncoder(writer)
	encoder.Encode(tasks)

}
