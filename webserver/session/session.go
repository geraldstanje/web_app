package session

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
  "log"
  "errors"
)

var cookieStore = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"user": userName,
	}
	if encoded, err := cookieStore.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Expires: time.Now().UTC().AddDate(1, 0, 0),
			Name:    "session",
			Value:   encoded,
			Path:    "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   "session",
		Path:   "/",
	}
	http.SetCookie(w, cookie)
}

func GetSessionUser(r *http.Request) (userName string, err error) {
	if cookie, err := r.Cookie("session"); err == nil {

    log.Println("GetSessionUser - cookie:", cookie.String())

    if cookie.MaxAge == -1 {
      err = errors.New("cookie expired")
    }

		cookieValue := make(map[string]string)
		if err = cookieStore.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["user"]
		}
	}
	return userName, err
}