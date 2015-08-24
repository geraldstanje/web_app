package authentication

import (
  "encoding/json"
  "log"
	d "github.com/geraldstanje/web_app/webserver/db"
	s "github.com/geraldstanje/web_app/webserver/session"
	"net/http"
)

type Message struct {
    Succeed bool `json:"succeed"`
    Error string `json:"error"`
    Redirect string `json:"redirect"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("email")
	pass := r.FormValue("password")
  succeed := false
  error := ""
  redirectTarget := "/"

	if user != "" && pass != "" && d.IsValidLogin(user, pass) {
		s.SetSession(user, w)
    succeed = true
    error = "Invalid user or password"
    redirectTarget = "/musicalbums"
    //http.Redirect(w, r, redirectTarget, 302)
	} //else {
    //str := "Invalid username or password"
    //w.Write([]byte(str))
	  //http.Redirect(w, r, redirectTarget, 302)
  //}

  m := Message{succeed, error, redirectTarget}

  //w.Header().Set("Content-Type", "application/json")
  //json.NewEncoder(w).Encode(m)
  if err := json.NewEncoder(w).Encode(m); err != nil {
    log.Println("[webserver] " + err.Error())
  }
}

func Logout(w http.ResponseWriter, r *http.Request) {
	s.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
