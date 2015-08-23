package authentication

import (
	d "github.com/geraldstanje/web_app/webserver/db"
	s "github.com/geraldstanje/web_app/webserver/session"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request) {
	user := request.FormValue("username")
	pass := request.FormValue("password")
	redirectTarget := "/"

	if user != "" && pass != "" {
		if !d.IsValidLogin(user, pass) {
			http.Redirect(response, request, redirectTarget, 302)
			return
		}
		s.SetSession(user, response)
		redirectTarget = "/musicalbums"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	s.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}
