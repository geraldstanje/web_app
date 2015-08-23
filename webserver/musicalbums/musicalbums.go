package musicalbums

import (
	"fmt"
	s "github.com/geraldstanje/web_app/webserver/session"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var size = "10"

type User struct {
	Username string
}

type Image struct {
  Filename  string
  Width   string
  Height string
}

const imageLink = `<img border=\"5\" style=\"margin:5px 5px\" src=\"{{.Filename}}\" width=\"{{.Width}}\" height=\"{{.Height}}\">`

func renderImgTemplate(filename string, width string, height string) (string, error) {
  img := Image{Filename: filename, Width: width, Height: height}
  t, err := template.New("imagelink").Parse(imageLink)
  if err != nil {
    return "", err
  }
  var out string
  err = t.Execute(out, img)
  if err != nil {
    return "", err
  }
  return out, nil
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
			img, err := renderImgTemplate(f.Name(), size, size)
      if err != nil {
        log.Fatal(err)
      }
      //img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\""+"files/%s"+"\" width=\""+"%s"+"\" height=\""+"%s"+"\">", f.Name(), size, size)
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
			img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\""+"files/%s"+"\" width=\""+"%s"+"\" height=\""+"%s"+"\">", f.Name(), size, size)
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
