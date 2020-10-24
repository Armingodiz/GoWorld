package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",home)
	http.Handle("/pictures/",http.StripPrefix("/pictures",http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080",nil)
}
func home(w http.ResponseWriter , req *http.Request)  {
	tpl:= template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.Execute(w,nil)
	if err!=nil{
		log.Fatalln(err)
	}
}