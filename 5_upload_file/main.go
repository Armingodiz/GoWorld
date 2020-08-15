package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

type data struct {
	Name, LName, Text string
	Subm              bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8585", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	var d data
	if req.Method == http.MethodPost {
		n := req.FormValue("name")
		l := req.FormValue("lname")
		s := req.FormValue("submitation") == "on"
		f, h, err := req.FormFile("fname")
		defer f.Close()
		handleErr(err)
		fmt.Println(f, h)
		reader, err2 := ioutil.ReadAll(f)
		handleErr(err2)
		d = data{n, l, string(reader), s}
		if s {
			nf, er := os.Create(filepath.Join("./newFiles/", h.Filename+"2"))
			handleErr(er)
			defer nf.Close()
			_, errr := nf.WriteString(string(reader))
			handleErr(errr)
		}
	} else {
		d = data{"", "", "", false}
	}
	err3 := tpl.ExecuteTemplate(w, "index.gohtml", d)
	handleErr(err3)
}
func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
