package authentication

import (
	"encoding/json"
	d "github.com/geraldstanje/web_app/webserver/db"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func FakeRegister(user string, pass string) {
  _ = d.RemoveUser(user)

  v := url.Values{}
  v.Add("email", user)
  v.Add("password", pass)

  req, _ := http.NewRequest("POST", "", strings.NewReader(v.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  w := httptest.NewRecorder()
  Register(w, req)

  if w.Code != http.StatusOK {
    t.Errorf("Home page didn't return %v", http.StatusOK)
  }

  var m Message
  if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
    t.Errorf(err.Error())
  }

  if m.Succeed != true {
    t.Errorf("Wrong json response")
  }
}

func TestRegister(t *testing.T) {
	FakeRegister("test@gmail.com", "root")
}

func TestLogin(t *testing.T) {
  FakeRegister("test@gmail.com", "root")

  v := url.Values{}
  v.Add("email", "test@gmail.com")
  v.Add("password", "root")

  req, _ := http.NewRequest("POST", "", strings.NewReader(v.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  w := httptest.NewRecorder()
  Login(w, req)

  if w.Code != http.StatusOK {
    t.Errorf("Home page didn't return %v", http.StatusOK)
  }

  var m Message
  if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
    t.Errorf(err.Error())
  }

  if m.Succeed != true || m.Redirect != "/musicalbums" {
    t.Errorf("Wrong json response")
  }
}