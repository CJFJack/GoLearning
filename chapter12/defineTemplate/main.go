package main

import (
	"os"
	"text/template"
)

func main() {
	tplText := `{{define "len"}} {{len .}} {{end}}
{{define "raw"}} {{.}} {{end}}
{{template "raw"}}
`

	tpl := template.Must(template.New("tpl").Parse(tplText))
	tpl.ExecuteTemplate(os.Stdout, "len", "123456")
	tpl.ExecuteTemplate(os.Stdout, "raw", "abcdef")
	tpl.ExecuteTemplate(os.Stdout, "raw", "qwerty")
}
