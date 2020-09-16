package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"tawesoft.co.uk/go/dialog"
)

type user struct {
	Name, Email, password string
}
type session struct {
	value string
}

var tpl *template.Template
var users = map[string]user{}
var sessions = map[string]session{} // cookie.value --> session

func init() {
	db, err := sql.Open("mysql",
		"root:armin3011@tcp(127.0.0.1:3306)/loginSystem")
	errHandler(err)
	err = db.Ping()
	errHandler(err)
	rows, err2 := db.Query("SELECT * FROM loginSystem.users")
	errHandler(err2)
	var x int
	var n, e, p string
	for rows.Next() {
		fmt.Println("###############################################################")
		fmt.Printf("var1 = %T\n", db)
		err := rows.Scan(&x, &n, &e, &p)
		errHandler(err)
		fmt.Println(n, e, p)
		new_user := user{n, e, p}
		users[e] = new_user
		fmt.Println(users)
	}
	defer db.Close()
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/signUp", signPage)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/logOut", logOut)

	http.Handle("/loginStatics/", http.StripPrefix("/loginStatics/", http.FileServer(http.Dir("Templates/loginStatics"))))
	http.Handle("/signStatics/", http.StripPrefix("/signStatics/", http.FileServer(http.Dir("Templates/signStatics"))))
	http.Handle("/mainStatics/", http.StripPrefix("/mainStatics/", http.FileServer(http.Dir("Templates/mainStatics"))))
	http.ListenAndServe(":8585", nil)
}
func loginPage(w http.ResponseWriter, req *http.Request) {
	if checkLog(req) {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var loggedUser user
	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		p := req.FormValue("pass")
		loggedUser, ok := users[e]
		if !ok {
			http.Error(w, "WRONG EMAIL ! ", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		if loggedUser.password != p {
			http.Error(w, "WRONG password ! ", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		sessions[c.Value] = session{e}
		http.Redirect(w, req, "/main", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "logPage.gohtml", loggedUser)
	errHandler(err)
}

func signPage(w http.ResponseWriter, req *http.Request) {
	if checkLog(req) {
		dialog.Alert("you are already logged in ! ")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var newUser user
	if req.Method == http.MethodPost {
		n := req.FormValue("name")
		e := req.FormValue("email")
		p := req.FormValue("password")
		fmt.Println(n, e, p)
		if _, ok := users[e]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			//dialog.Alert("THERE IS ALREADY AN ACCOUNT WITH THIS EMAIL ! ")
			return
		}
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		sessions[cookie.Value] = session{e}
		newUser = user{n, e, p}
		users[e] = newUser
		db, err := sql.Open("mysql",
			"root:armin3011@tcp(127.0.0.1:3306)/loginSystem")
		errHandler(err)
		stmt, err := db.Prepare(`INSERT INTO users VALUES (?,?,?,?);`)
		defer stmt.Close()
		defer db.Close()
		r, err := stmt.Exec(10, n, e, p)
		errHandler(err)
		ro, err := r.RowsAffected()
		errHandler(err)
		fmt.Println("INSERTED RECORD", ro)
		errHandler(err)
		http.Redirect(w, req, "/main", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "sign.gohtml", newUser)
	errHandler(err)
}
func logOut(w http.ResponseWriter, req *http.Request) {
	if !checkLog(req) {
		dialog.Alert("YOU ARE NOT LOGGED IN YET ! ")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	c, err := req.Cookie("session")
	errHandler(err)
	delete(sessions, c.Value)
	cook := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cook)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func mainPage(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if u.Name != "" {
		err := tpl.ExecuteTemplate(w, "index.gohtml", u)
		errHandler(err)
		return
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	errHandler(err)
}
func errHandler(er error) {
	if er != nil {
		log.Fatalln(er)
	}
}
