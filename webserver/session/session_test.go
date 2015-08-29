package session

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "log"
)

func TestSetSession(t *testing.T) {
  w := httptest.NewRecorder()
  SetSession("Douglas.Costa@gmail.com", w)

  req, _ := http.NewRequest("POST", "", nil)
  c, err := w.Cookie("Douglas.Costa@gmail.com")
  if err != nil {
    t.Errorf("get Cookie failed")
  }

  req.AddCookie(c)
  user := GetUserName(req)
  log.Println("User:", user)

  //if user != "Douglas.Costa@gmail.com" {
  //  t.Errorf("GetUserName failed")
  //}
}