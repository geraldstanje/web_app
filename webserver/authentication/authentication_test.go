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

func TestRegister(t *testing.T) {
	t.Log("Run TestRegister...")

	_ = d.RemoveUser("test@gmail.com")

	v := url.Values{}
	v.Add("email", "test@gmail.com")
	v.Add("password", "root")

	req, _ := http.NewRequest("POST", "", strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	Register(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	var m Message
	if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
		t.Errorf(err.Error())
	}

	t.Log("m: ", m)

	t.Log("Run TestRegister...successful")
}
