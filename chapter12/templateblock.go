package main

import (
	"os"
	"text/template"
)

func main() {
	tplText := `输入内容：{{block "content" .}} {{.}} {{end}}`
	tpl := template.Must(template.New("tpl").Parse(tplText))
	tpl2, _ := template.Must(tpl.Clone()).Parse(`{{define "content"}}{{.}}{{end}}`)

	tpl.Execute(os.Stdout, "abcdefg")
	tpl2.Execute(os.Stdout, "abcdefg")
}
