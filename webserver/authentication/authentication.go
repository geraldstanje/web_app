package authentication

import (
  "net/http"
  s "github.com/geraldstanje/web_app/webserver/session"
  d "github.com/geraldstanje/web_app/webserver/db"
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
    s.SetSession(name, response)
    redirectTarget = "/musicalbums"
  }
  http.Redirect(response, request, redirectTarget, 302)
}

func Logout(response http.ResponseWriter, request *http.Request) {
  s.ClearSession(response)
  http.Redirect(response, request, "/", 302)
}