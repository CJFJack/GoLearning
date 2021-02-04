package main

import (
	"bufio"
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
	tasks := []Task{}
	file, _ := os.Open("task2.inputJson")
	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)
	decoder.Decode(&tasks)

	fmt.Printf("%#v", tasks)
}
