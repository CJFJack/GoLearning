package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	tplText := `{{range.}}
{{.Id}}-{{.Name}}-{{if eq .Sex 1}}男{{else}}女{{end}}
{{end}}`
	tpl := template.Must(template.New("tpl").Parse(tplText))
	//tpl.Execute(os.Stdout, "<h1>我是标题</h>")
	//tpl.Execute(os.Stdout, []int{1,2,3,4})
	//tpl.Execute(os.Stdout, map[string]int{"a":1,"b":2})
	//tpl.Execute(os.Stdout, struct{
	//	Id int
	//	Name string
	//	Sex int
	//}{Id:1, Name:"cjf", Sex:1})
	tpl.Execute(os.Stdout, []struct {
		Id   int
		Name string
		Sex  int
	}{{Id: 1, Name: "cjf", Sex: 1}, {Id: 2, Name: "lulu", Sex: 0}})

	s := "abc"
	fmt.Printf("%s%s\n", strings.ToUpper(s[:1]), s[1:])
	fmt.Printf("%T", s[0])

}
