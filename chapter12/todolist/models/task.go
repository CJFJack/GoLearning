package models

type Task struct {
	ID     int
	Name   string
	Status string
}

var tasks = []*Task{&Task{ID: 1, Name: "洗衣服", Status: "正在执行"}, &Task{ID: 2, Name: "做作业", Status: "未执行"}, &Task{ID: 3, Name: "写代码", Status: "未执行"}}

func GetTasks() []*Task {
	return tasks
}

func AddTasks(name string) {
	tasks = append(tasks, &Task{len(tasks) + 1, name, "新增"})
}
