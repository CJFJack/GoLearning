package main

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	toDoList = make([]map[string]string, 0)
	text     string
)

const (
	passSalt   string = "61info"
	saltedPass string = "1284aff560103c123416967d2f3ac290"
)

func input(inputText string) string {
	fmt.Print(inputText)
	fmt.Scan(&text)
	return text
}

func genId() int {
	var maxId int

	if len(toDoList) == 0 {
		return 1
	} else {
		for _, value := range toDoList {
			id, _ := strconv.Atoi(value["id"])
			if id > maxId {
				maxId = id
			}
		}
		return maxId + 1
	}
}

func add() {

	name := input("请输入任务名称：")
	startTime := input("请输入开始时间:")
	user := input("请输入负责人：")

	id := strconv.Itoa(genId())

	toDoList = append(toDoList, map[string]string{"id": id, "name": name, "startTime": startTime, "user": user, "endTime": "", "status": "新创建"})

}

func printToDo(t map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("任务id：", t["id"])
	fmt.Println("任务名称：", t["name"])
	fmt.Println("开始时间：", t["startTime"])
	fmt.Println("结束时间：", t["endTime"])
	fmt.Println("负责人：", t["user"])
	fmt.Println("状态：", t["status"])
	fmt.Println(strings.Repeat("-", 20))
}

func query() {
	name := input("请输入需要查询的任务名称，输入all则查询所有：")
	n := 0
	for _, task := range toDoList {
		if task["name"] == name || name == "all" {
			printToDo(task)
			n++
		}
	}
	if n == 0 {
		fmt.Println("任务不存在")
	}
}

func sortTask() {
	sortTypeArray := [2]string{"name", "startTime"}
	n := 0
	sortType := input("请输入排序方式，name:按任务名称升序，time:按开始时间升序，否则按任务id升序：")
	for _, value := range sortTypeArray {
		if sortType == value {
			sort.SliceStable(toDoList, func(i, j int) bool {
				return toDoList[i][sortType] < toDoList[j][sortType]
			})
			n++
		}
	}
	if n == 0 {
		sort.SliceStable(toDoList, func(i, j int) bool {
			return toDoList[i]["id"] < toDoList[j]["id"]
		})
	}
	createTable()
}

func delToDoListTask(index int) {
	copy(toDoList[index:], toDoList[index+1:])
	toDoList = toDoList[:len(toDoList)-1]
}

func del() {
	name := input("请输入要删除的任务名称：")
	for index, task := range toDoList {
		if task["name"] == name {
			delToDoListTask(index)
			break
		}
	}
}

func modify() {
	name := input("请输入需要修改的任务：")
	startTime := input("请输入开始时间：")
	user := input("请输入负责人：")
	status := input("请输入任务状态(新创建/进行中/已完成)：")
	for _, task := range toDoList {
		if task["name"] == name {
			task["startTime"] = startTime
			task["user"] = user
			task["status"] = status
			if status == "已完成" {
				task["endTime"] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
		break
	}
}

func actionTransfer(action string) func() {
	switch {
	case action == "add":
		return add
	case action == "query":
		return query
	case action == "del":
		return del
	case action == "modify":
		return modify
	case action == "sort":
		return sortTask
	default:
		fmt.Println("指令不正确")
		return func() {

		}
	}
}

func calcSaltMd5(rawPass string) string {
	md5Pass := fmt.Sprintf("%x", md5.Sum([]byte(rawPass+passSalt)))
	return string(md5Pass)
}

func createTable() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"任务Id", "任务名称", "开始时间", "结束时间", "负责人", "状态"})
	for _, toDoValue := range toDoList {
		row := make([]string, 0)
		row = append(row, toDoValue["id"], toDoValue["name"], toDoValue["startTime"], toDoValue["endTime"], toDoValue["user"], toDoValue["status"])
		t.Append(row)
	}
	t.Render()
}

func main() {
	fmt.Print("请输入密码：")
	rawPass, _ := gopass.GetPasswd()
	inputSaltedPass := calcSaltMd5(string(rawPass))
	if inputSaltedPass != saltedPass {
		fmt.Println("您输入的密码错误，退出")
		return
	}
	for {
		action := input("请输入query/sort/add/modify/del/exit：")
		if action == "exit" {
			break
		}
		actionTransfer(action)()
	}
}
