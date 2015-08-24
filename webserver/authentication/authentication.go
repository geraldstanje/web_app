package authentication

import (
	"encoding/json"
	d "github.com/geraldstanje/web_app/webserver/db"
	s "github.com/geraldstanje/web_app/webserver/session"
	"log"
	"net/http"
)

type Message struct {
	Succeed  bool   `json:"succeed"`
	Info     string `json:"info"`
	Redirect string `json:"redirect"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("email")
	pass := r.FormValue("password")
	succeed := false
	info := ""

	if user != "" && pass != "" && d.IsValidRegistration(user, pass) {
		succeed = true
    info = "Registration succeeded"
	} else {
		info = "Registration failed, user already registered"
	}

	m := Message{succeed, info, ""}

	//w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Println("[webserver] " + err.Error())
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("email")
	pass := r.FormValue("password")
	succeed := false
	info := ""
	redirectTarget := "/"

	if user != "" && pass != "" && d.IsValidLogin(user, pass) {
		s.SetSession(user, w)
		succeed = true
		redirectTarget = "/musicalbums"
	} else {
		info = "Invalid user or password"
	}

	m := Message{succeed, info, redirectTarget}

	//w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Println("[webserver] " + err.Error())
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	s.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
