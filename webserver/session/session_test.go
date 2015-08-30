package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

// Suggestion: Make the encoded value setter/getter function more abstract.

func TestSetSession(t *testing.T) {
	w := httptest.NewRecorder()
	if err := SetSession("test", w); err != nil {
		t.Errorf("SetSession error: %s", err)
	}

	c, err := getRecordedCookie(w, "session")
	if err != nil {
		t.Fatalf("getRecordedCookie error: %s", err)
	}
	if c.MaxAge <= 0 {
		t.Error("cookie not set properly")
	}
}

func TestGetSession(t *testing.T) {
	value := map[string]string{
		"user": "test",
	}
	encoded, err := cookieStore.Encode("session", value)
	if err != nil {
		t.Fatal("encoding failed")
	}

	c := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    "session",
		Value:   encoded,
		Path:    "/",
	}
	req, _ := http.NewRequest("GET", "", nil)
	req.AddCookie(c)

	user, err := GetSessionUser(req)
	if err != nil {
		t.Errorf("GetSessionUser error: %s", err)
	}
	if user != "test" {
		t.Errorf("expected user %q, got user %q", "test", user)
	}
}

func TestClearSession(t *testing.T) {
	w := httptest.NewRecorder()
	ClearSession(w)

	c, err := getRecordedCookie(w, "session")
	if err != nil {
		t.Fatalf("getRecordedCookie error: %s", err)
	}
	if c.MaxAge > 0 {
		t.Error("cookie will not be exterminated")
	}
}
