package session

import (
  "testing"
  "net/http"
  "net/http/httptest"
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

  req, _ := http.NewRequest("POST", "", nil)
  req.AddCookie(c)
  user := GetUserName(req)

  if user != "Douglas.Costa@gmail.com" {
    t.Errorf("GetUserName failed")
  }

  ClearSession(w)
  c, err = getRecordedCookie(w, "session")
  if err != nil {
    t.Errorf("getRecordedCookie failed")
  }
  req.AddCookie(c)

  user = GetUserName(req)

  log.Println("User", user)

  //if user != "" {
  //  t.Errorf("GetUserName failed")
  //}
}