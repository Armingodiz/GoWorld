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
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(f, h)
		reader, err2 := ioutil.ReadAll(f)
		if err2 != nil {
			log.Fatalln(err2)
		}
		d = data{n, l, string(reader), s}
		if s {
			nf, er := os.Create(filepath.Join("./newFiles/", h.Filename+"2"))
			if er != nil {
				log.Fatalln(er)
			}
			defer nf.Close()
			_, errr := nf.WriteString(string(reader))
			if errr != nil {
				log.Fatalln(errr)
			}
		}
	} else {
		d = data{"", "", "", false}
		fmt.Println("FUCK ME")
	}
	err3 := tpl.ExecuteTemplate(w, "index.gohtml", d)
	if err3 != nil {
		log.Println(err3)
	}
}
