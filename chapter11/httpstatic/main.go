package main

import "net/http"

func main() {
	//http.Handle("/www/", http.FileServer(http.Dir(".")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./www"))))

	http.ListenAndServe(":8888", nil)
}
