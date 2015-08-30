package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "log"
)

func getRecordedCookie(recorder *httptest.ResponseRecorder, name string) (*http.Cookie, error) {
	r := &http.Response{Header: recorder.HeaderMap}
	for _, cookie := range r.Cookies() {
		if cookie.Name == name {
			return cookie, nil
		}
	}
	return nil, http.ErrNoCookie
}

func TestSetSession(t *testing.T) {
	w := httptest.NewRecorder()
	SetSession("Douglas.Costa@gmail.com", w)

	c, err := getRecordedCookie(w, "session")
	if err != nil {
		t.Errorf("getRecordedCookie failed")
	}

	req, _ := http.NewRequest("GET", "", nil)
	req.AddCookie(c)
	user, err := GetSessionUser(req)
	if err != nil {
		t.Errorf("GetSessionUser failed")
	}

	if user != "Douglas.Costa@gmail.com" {
		t.Errorf("GetSessionUser failed")
	}
}

func TestClearSession(t *testing.T) {
  w := httptest.NewRecorder()
  SetSession("David.Alaba@gmail.com", w)

  c, err := getRecordedCookie(w, "session")
  if err != nil {
    t.Errorf("getRecordedCookie failed")
  }

  req, _ := http.NewRequest("GET", "", nil)
  req.AddCookie(c)
  user, err := GetSessionUser(req)
  if err != nil {
    t.Errorf("GetSessionUser failed")
  }

  if user != "David.Alaba@gmail.com" {
    t.Errorf("GetSessionUser failed")
  }

  w.Header().Del("Set-Cookie")
  ClearSession(w)

  c, err = getRecordedCookie(w, "session")

  log.Println("cookie:", c.String())

  if err != nil {
    t.Errorf("getRecordedCookie failed")
  }

	req, _ = http.NewRequest("GET", "", nil)
  req.AddCookie(c)

	user, err = GetSessionUser(req)
	if err == nil {
		t.Errorf("GetSessionUser failed")
	}
}