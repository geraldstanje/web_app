package authentication

import (
  "net/http"
  s "github.com/geraldstanje/web_app/webserver/session"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
  name := request.FormValue("name")
  pass := request.FormValue("password")
  redirectTarget := "/"
  if name != "" && pass != "" {
    // .. check credentials ..
    s.SetSession(name, response)
    redirectTarget = "/musicalbums"
  }
  http.Redirect(response, request, redirectTarget, 302)
}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
  s.ClearSession(response)
  http.Redirect(response, request, "/", 302)
}