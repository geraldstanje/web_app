package session

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "log"
)

func TestSetSession(t *testing.T) {
  req, _ := http.NewRequest("POST", "", nil)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  w := httptest.NewRecorder()

  SetSession("Douglas.Costa@gmail.com", w)

  user := GetUserName(req)
  
  log.Println("User:", user)

  //if user != "Douglas.Costa@gmail.com" {
  //  t.Errorf("GetUserName failed")
  //}
}