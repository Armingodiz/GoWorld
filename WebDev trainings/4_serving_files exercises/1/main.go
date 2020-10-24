package main

import (
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "WELL COME TO HOME PAGE ! ")
}
func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}
func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
