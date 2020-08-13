package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/", http.HandlerFunc(goIndex))
	http.Handle("/about", http.HandlerFunc(goAbout))
	http.Handle("/sign", http.HandlerFunc(goSign))
	http.ListenAndServe(":8080", nil)
}

func goIndex(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	HandleError(res, err)
}

func goAbout(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req.Form)
		err2 := tpl.ExecuteTemplate(res, "about.gohtml", req.Form)
		HandleError(res, err2)
	case "GET":
		err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
		HandleError(res, err)
	default:
		break
	}

}

func goSign(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "submit.gohtml", nil)
	HandleError(res, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
