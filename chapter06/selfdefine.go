package main

import "fmt"

type Counter int
type Task map[string]string
type Callback func(...string)

func main() {
	var cnt Counter
	cnt = 1
	fmt.Printf("%T, %#v\n", cnt, cnt)

	var task Task
	task = map[string]string{"name": "cjf"}
	task["age"] = "30"
	fmt.Printf("%T, %#v\n", task, task)

	var callback Callback
	callback = func(s ...string) {
		for i, v := range s {
			fmt.Println(i, v)
		}
	}
	callback("a", "b", "c")

}
