package main

import(
	"io"
	"net/http"
)

func homePage(w http.ResponseWriter , req *http.Request)  {
	io.WriteString(w,"this is home page")
}
func about(w http.ResponseWriter , req *http.Request)  {
	io.WriteString(w,"this is about page")
}
func armin(w http.ResponseWriter , req *http.Request)  {
	io.WriteString(w,"Hello user Armin")
}
func main() {
	http.HandleFunc("/",homePage)
	http.HandleFunc("/about",about)
	http.HandleFunc("/me/",armin)
	http.ListenAndServe(":8585",nil)
}
