ackage session

import (
  "github.com/gorilla/securecookie"
  "net/http"
  "strconv"
  "time"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(userName string, w http.ResponseWriter) {
  t := strconv.FormatInt(time.Now().UnixNano(), 10)
  value := map[string]string{
    "user": userName,
    "time": t,
  }
  if encoded, err := cookieHandler.Encode("session", value); err == nil {
    cookie := &http.Cookie{Name: "session", Value: encoded, Path: "/"}
    http.SetCookie(w, cookie)
  }
}

func ClearSession(w http.ResponseWriter) {
  t := strconv.FormatInt(time.Now().UnixNano(), 10)
  value := map[string]string{
    "time": t,
  }
  if encoded, err := cookieHandler.Encode("session", value); err == nil {
    cookie := &http.Cookie{Name: "session", Value: encoded, MaxAge: -1, Expires: time.Unix(1, 0)}
    http.SetCookie(w, cookie)
  }
}

func GetSessionUser(r *http.Request) (userName string, err error) {
  if cookie, err := r.Cookie("session"); err == nil {
    cookieValue := make(map[string]string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      userName = cookieValue["user"]
    }
  }
  return userName, err
}

func GetSessionTime(r *http.Request) (t string, err error) {
  if cookie, err := r.Cookie("session"); err == nil {
    cookieValue := make(map[string]string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      t = cookieValue["time"]
    }
  }
  return t, err
}