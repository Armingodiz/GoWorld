package main

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user
	cookie, err := req.Cookie("session")

	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	http.SetCookie(w, cookie)
	if ses, er := sessions[cookie.Value]; er {
		u, _ = users[ses.value]
	}
	return u
}

func checkLog(req *http.Request) bool {
	cook, er := req.Cookie("session")
	if er != nil {
		return false
	}
	sess := sessions[cook.Value]
	_, ok := users[sess.value]
	return ok
}
