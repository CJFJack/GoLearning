package main

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	storeType = "inputJson"
)

const (
	passSalt   = "61info"
	jsonPrefix = "tasks.inputJson."
	gobPrefix  = "tasks.gob."
)

type Tasks struct {
	Id        int
	Name      string
	StartTime string
	User      string
	EndTime   string
	Status    string
}

type Store interface {
	Encode([]*Tasks) error
	Decode(*[]*Tasks) error
}

type Json struct{}

func (j *Json) Encode(tasks []*Tasks) error {
	saveFile := CreateOrDeleteSaveFile(jsonPrefix)
	file, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	encoder := json.NewEncoder(writer)
	return encoder.Encode(tasks)
}

func (j *Json) Decode(tasks *[]*Tasks) error {
	maxSuffix, _ := MaxSuffixFileNameAndClearOldFile(jsonPrefix, false)
	file, err := os.OpenFile(jsonPrefix+maxSuffix, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)
	return decoder.Decode(tasks)
}

type Gob struct{}

func (g *Gob) Encode(tasks []*Tasks) error {
	saveFile := CreateOrDeleteSaveFile(gobPrefix)
	file, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	encoder := gob.NewEncoder(writer)
	return encoder.Encode(tasks)
}

func (g *Gob) Decode(tasks *[]*Tasks) error {
	maxSuffix, _ := MaxSuffixFileNameAndClearOldFile(gobPrefix, false)
	file, err := os.OpenFile(gobPrefix+maxSuffix, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := gob.NewDecoder(reader)
	return decoder.Decode(tasks)
}

var (
	toDoList = []*Tasks{}
	text     string
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
			id := task.Id
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

	task := Tasks{Id: id, Name: name, StartTime: startTime, User: user, Status: "新创建"}
	toDoList = append(toDoList, &task)

	StoreEncode()

}

func printToDo(t *Tasks) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("任务id：", t.Id)
	fmt.Println("任务名称：", t.Name)
	fmt.Println("开始时间：", t.StartTime)
	fmt.Println("结束时间：", t.EndTime)
	fmt.Println("负责人：", t.User)
	fmt.Println("状态：", t.Status)
	fmt.Println(strings.Repeat("-", 20))
}

func query() {
	name := input("请输入需要查询的任务名称，输入all则查询所有：")
	n := 0
	for _, task := range toDoList {
		if task.Name == name || name == "all" {
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
	sortType := input("请输入排序方式，name:按任务名称升序，startTime:按开始时间升序，否则按任务id升序：")
	for _, value := range sortTypeArray {
		if sortType == value {
			sort.SliceStable(toDoList, func(i, j int) bool {
				if sortType == "name" {
					return toDoList[i].Name < toDoList[j].Name
				} else if sortType == "startTime" {
					return toDoList[i].StartTime < toDoList[j].StartTime
				} else {
					return true
				}
			})
			n++
		}
	}
	if n == 0 {
		sort.SliceStable(toDoList, func(i, j int) bool {
			return toDoList[i].Id < toDoList[j].Id
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
		if task.Name == name {
			delToDoListTask(index)
			break
		}
	}

	StoreEncode()
}

func modify() {
	name := input("请输入需要修改的任务：")
	startTime := input("请输入开始时间：")
	//endTime := input("请输入结束时间：")
	user := input("请输入负责人：")
	status := input("请输入任务状态(新创建/进行中/已完成)：")
	for _, task := range toDoList {
		if task.Name == name {
			fmt.Println(task)
			task.StartTime = startTime
			//task.endTime = endTime
			task.User = user
			task.Status = status
			if status == "已完成" {
				task.EndTime = time.Now().Format("2006-01-02 15:04:05")
			}
			break
		}
	}

	StoreEncode()
}

func storeSelect() {
	storeType = input("请选择持久化类型json/gob：")
	err := StoreDecode(storeType)
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
	}
}

func actionTransfer(action string) func() {
	// 方法一
	actionMap := map[string]func(){
		"add":    add,
		"query":  query,
		"del":    del,
		"modify": modify,
		"sort":   sortTask,
		"store":  storeSelect,
	}

	if actionMap[action] == nil {
		fmt.Println("指令不不正确")
		return func() {

		}
	}
	return actionMap[action]

	// 方法二
	//switch {
	//case action == "add":
	//	return add
	//case action == "query":
	//	return query
	//case action == "del":
	//	return del
	//case action == "modify":
	//	return modify
	//case action == "sort":
	//	return sortTask
	//default:
	//	fmt.Println("指令不正确")
	//	return func() {
	//
	//	}
	//}
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
		row = append(row, strconv.Itoa(task.Id), task.Name, task.StartTime, task.EndTime, task.User, task.Status)
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

func MaxSuffixFileNameAndClearOldFile(prefix string, clear bool) (string, string) {
	files, _ := filepath.Glob(prefix + "*")
	maxSuffix := "0"
	newSuffix := "0"
	if len(files) > 0 {
		sort.Strings(files)
		fileSplit := strings.Split(files[len(files)-1], prefix)
		maxSuffix = fileSplit[len(fileSplit)-1]
		maxId, _ := strconv.Atoi(maxSuffix)
		if clear && len(files) >= 3 {
			for i := 0; i <= len(files)-3; i++ {
				os.Remove(files[i])
			}
		}
		newSuffix = strconv.Itoa(maxId + 1)
	}
	return maxSuffix, newSuffix
}

func CreateOrDeleteSaveFile(prefix string) string {
	_, newSuffix := MaxSuffixFileNameAndClearOldFile(prefix, true)
	return prefix + newSuffix
}

func StoreDecode(storeType string) error {
	switch storeType {
	case "inputJson":
		var json Store = &Json{}
		return json.Decode(&toDoList)
	case "gob":
		var gob Store = &Gob{}
		return gob.Decode(&toDoList)
	default:
		return errors.New("持久化类型不正确")
	}
}

func StoreEncode() error {
	switch storeType {
	case "inputJson":
		var json Store = &Json{}
		return json.Encode(toDoList)
	case "gob":
		var gob Store = &Gob{}
		return gob.Encode(toDoList)
	default:
		return errors.New("持久化类型不正确")
	}
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

	err := StoreDecode(storeType)
	if err != nil {
		if err != io.EOF {
			fmt.Println(err)
			return
		}
	}

	for {
		action := input("请输入store/query/sort/add/modify/del/exit：")
		if action == "exit" {
			break
		}
		actionTransfer(action)()
	}

}
