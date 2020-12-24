package main

import (
	"html/template"
	"net/http"
	"todolist/models"
)

func main() {
	addr := ":9999"

	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseGlob("views/list.html"))
		tpl.ExecuteTemplate(writer, "list.html", models.GetTasks())
	})

	http.HandleFunc("/add/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			tpl := template.Must(template.ParseFiles("views/add.html"))
			tpl.ExecuteTemplate(writer, "add.html", nil)
		}
		if request.Method == "POST" {
			name := request.PostFormValue("name")
			models.AddTasks(name)
			http.Redirect(writer, request, "/list/", 302)
		}

	})

	http.ListenAndServe(addr, nil)
}
