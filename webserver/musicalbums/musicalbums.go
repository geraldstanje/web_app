package musicalbums

import (
  s "github.com/geraldstanje/web_app/webserver/session"
  "io/ioutil"
  "net/http"
  "log"
  "fmt"
  "html/template"
)

var size = "10"

type User struct {
  Username string
}

func Upload(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    _, header, err := r.FormFile("TheFile")
    if err != nil {
      log.Fatal(err)
    }
    file, _ := header.Open()
    path := fmt.Sprintf("files/%s", header.Filename)
    buf, _ := ioutil.ReadAll(file)
    ioutil.WriteFile(path, buf, 0644)

    files, _ := ioutil.ReadDir("./files")
    for _, f := range files {
      img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\"" + "files/%s" + "\" width=\"" + "%s" + "\" height=\"" + "%s" + "\">", f.Name(), size, size)
      w.Write([]byte(img))
    }
  }
 }

func Resize(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    size = r.FormValue("value")
    if size == "" {
      fmt.Println("Empty FormValue")
      return
    }

    files, _ := ioutil.ReadDir("./files")
    for _, f := range files {
      img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\"" + "files/%s" + "\" width=\"" + "%s" + "\" height=\"" + "%s" + "\">", f.Name(), size, size)
      w.Write([]byte(img))
    }
  }
}

func MusicAlbums(w http.ResponseWriter, req *http.Request) {
  username := s.GetUserName(req)
  if username != "" {
    user := User{Username: username}
    musicAlbumsTemplate, err := template.ParseFiles("templates/musicalbums.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err = musicAlbumsTemplate.Execute(w, user)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  } else {
    http.Redirect(w, req, "/", 302)
  }
}