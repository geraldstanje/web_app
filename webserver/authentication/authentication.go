package authentication

import (
	d "github.com/geraldstanje/web_app/webserver/db"
	s "github.com/geraldstanje/web_app/webserver/session"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("username")
	pass := r.FormValue("password")
	redirectTarget := "/"

	if user != "" && pass != "" && d.IsValidLogin(user, pass) {
		s.SetSession(user, w)
		redirectTarget = "/musicalbums"
  }
	http.Redirect(w, r, redirectTarget, 302)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	s.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}