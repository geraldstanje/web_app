package session

import (
	"github.com/gorilla/securecookie"
	"net/http"
  "errors"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"user": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func GetSessionUser(r *http.Request) (userName string, err error) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      if cookie.MaxAge == -1 {
        err = errors.New("session expired")
      } else {
			  userName = cookieValue["user"]
      }
		}
	}
	return userName, err
}