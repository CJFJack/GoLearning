package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Tasks struct {
	id        int
	name      string
	startTime string
	user      string
	endTime   string
	status    string
}

var (
	toDoList = make([]*Tasks, 0)
	text     string
)

const (
	passSalt string = "61info"
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
		for _, task := range toDoList {
			id := task.id
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

	id := genId()

	task := &Tasks{id: id, name: name, startTime: startTime, user: user, status: "新创建"}
	toDoList = append(toDoList, task)

	jsonEncode(toDoList)

}

func printToDo(t Tasks) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("任务id：", t.id)
	fmt.Println("任务名称：", t.name)
	fmt.Println("开始时间：", t.startTime)
	fmt.Println("结束时间：", t.endTime)
	fmt.Println("负责人：", t.user)
	fmt.Println("状态：", t.status)
	fmt.Println(strings.Repeat("-", 20))
}

func query() {
	name := input("请输入需要查询的任务名称，输入all则查询所有：")
	n := 0
	for _, task := range toDoList {
		if task.name == name || name == "all" {
			printToDo(*task)
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
	sortType := input("请输入排序方式，name:按任务名称升序，startTime:按开始时间升序，否则按任务id升序：")
	for _, value := range sortTypeArray {
		if sortType == value {
			sort.SliceStable(toDoList, func(i, j int) bool {
				if sortType == "name" {
					return toDoList[i].name < toDoList[j].name
				} else if sortType == "startTime" {
					return toDoList[i].startTime < toDoList[j].startTime
				} else {
					return true
				}
			})
			n++
		}
	}
	if n == 0 {
		sort.SliceStable(toDoList, func(i, j int) bool {
			return toDoList[i].id < toDoList[j].id
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
		if task.name == name {
			delToDoListTask(index)
			break
		}
	}
}

func modify() {
	name := input("请输入需要修改的任务：")
	startTime := input("请输入开始时间：")
	//endTime := input("请输入结束时间：")
	user := input("请输入负责人：")
	status := input("请输入任务状态(新创建/进行中/已完成)：")
	for _, task := range toDoList {
		if task.name == name {
			fmt.Println(task)
			task.startTime = startTime
			//task.endTime = endTime
			task.user = user
			task.status = status
			if status == "已完成" {
				task.endTime = time.Now().Format("2006-01-02 15:04:05")
			}
			break
		}
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
	for _, task := range toDoList {
		row := make([]string, 0)
		row = append(row, strconv.Itoa(task.id), task.name, task.startTime, task.endTime, task.user, task.status)
		t.Append(row)
	}
	t.Render()
}

func fileIsExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
	return true
}

func createPassFile(rawPass1 []byte) {
	passFile, err := os.Create("password")
	if err != nil {
		return
	}
	defer passFile.Close()
	inputSaltedPass := calcSaltMd5(string(rawPass1))
	passFile.Write([]byte(inputSaltedPass))
}

func readPassFile() string {
	saltedPass := ""
	passFile, err := os.Open("password")
	if err != nil {
		return saltedPass
	}
	defer passFile.Close()
	ctx := make([]byte, 10)
	for {
		n, readErr := passFile.Read(ctx)
		if readErr == io.EOF {
			break
		}
		saltedPass += string(ctx[:n])
	}
	return saltedPass
}

func jsonEncode(tasks []*Tasks) error {
	file, err := os.OpenFile("tasks.inputJson", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	encoder := json.NewEncoder(writer)
	return encoder.Encode(tasks)
}

func main() {
	exist := fileIsExist("password")
	if !exist {
	LABEL:
		fmt.Print("请输入初始化密码：")
		rawPass1, _ := gopass.GetPasswd()
		fmt.Print("请再次输入初始化密码：")
		rawPass2, _ := gopass.GetPasswd()
		if string(rawPass1) != string(rawPass2) {
			fmt.Println("两次密码输入一致")
			goto LABEL
		}
		createPassFile(rawPass1)

	} else {
		fmt.Print("请输入密码：")
		rawPass, _ := gopass.GetPasswd()
		inputSaltedPass := calcSaltMd5(string(rawPass))

		saltedPass := readPassFile()

		if inputSaltedPass != saltedPass {
			fmt.Println("您输入的密码错误，退出")
			return
		}
	}

	for {
		action := input("请输入query/sort/add/modify/del/exit：")
		if action == "exit" {
			break
		}
		actionTransfer(action)()
	}

}
