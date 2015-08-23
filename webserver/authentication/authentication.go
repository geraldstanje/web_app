package authentication

import (
  "net/http"
  s "github.com/geraldstanje/web_app/webserver/session"
)

func Login(response http.ResponseWriter, request *http.Request) {
  name := request.FormValue("username")
  pass := request.FormValue("password")
  redirectTarget := "/"
  if name != "" && pass != "" {
    // .. check credentials ..
    s.SetSession(name, response)
    redirectTarget = "/musicalbums"
  }
  http.Redirect(response, request, redirectTarget, 302)
}

func Logout(response http.ResponseWriter, request *http.Request) {
  s.ClearSession(response)
  http.Redirect(response, request, "/", 302)
}