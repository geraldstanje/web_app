package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var cookieStore = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(user string, w http.ResponseWriter) (err error) {
	val := map[string]string{
		"user": user,
	}
	if enc, err := cookieStore.Encode("session", val); err == nil {
		c := &http.Cookie{
			Expires: time.Now().UTC().AddDate(1, 0, 0),
			Name:    "session",
			Value:   enc,
			MaxAge:  3600,
			Path:    "/",
		}
		http.SetCookie(w, c)
	}
	return err
}

func ClearSession(w http.ResponseWriter) {
	c := &http.Cookie{Name: "session", MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, c)
}

func GetSessionUser(r *http.Request) (user string, err error) {
	if c, err := r.Cookie("session"); err == nil {
		if c.MaxAge <= 0 {
			err = errors.New("cookie expired")
		}

		val := make(map[string]string)
		if err = cookieStore.Decode("session", c.Value, &val); err == nil {
			user = val["user"]
		}
	}
	return user, err
}
