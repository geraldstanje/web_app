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

type UserAccount struct {
	User string
}

type Image struct {
	Filename string
	Width    string
	Height   string
}

const imageLink = `<img border="5" style="margin:5px 5px" src="files/{{.Filename}}" width="{{.Width}}" height="{{.Height}}">`

func renderImgTemplate(w http.ResponseWriter, filename string, width string, height string) error {
	img := Image{Filename: filename, Width: width, Height: height}
	t, err := template.New("imagelink").Parse(imageLink)
	if err != nil {
		return err
	}
	err = t.Execute(w, img)
	return err
}

func renderMusicAlbumsTemplate(w http.ResponseWriter, user string) error {
	useraccount := UserAccount{User: user}
	t, err := template.ParseFiles("templates/musicalbums.html")
	if err != nil {
		return err
	}
	err = t.Execute(w, useraccount)
	return err
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, err := r.FormFile("TheFile")
		if err != nil {
			log.Println("[webserver] " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file, _ := header.Open()
		path := fmt.Sprintf("files/%s", header.Filename)
		buf, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(path, buf, 0644)

		files, _ := ioutil.ReadDir("./files")
		for _, f := range files {
			err = renderImgTemplate(w, f.Name(), size, size)
			if err != nil {
				log.Println("[webserver] " + err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func Resize(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		size = r.FormValue("value")
		if size == "" {
			log.Println("[webserver] " + "Empty FormValue")
			http.Error(w, "Empty FormValue", http.StatusInternalServerError)
		}

		files, _ := ioutil.ReadDir("./files")
		for _, f := range files {
			err := renderImgTemplate(w, f.Name(), size, size)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func MusicAlbums(w http.ResponseWriter, req *http.Request) {
	user := s.GetUserName(req)
	if user != "" {
		err := renderMusicAlbumsTemplate(w, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, req, "/", 302)
	}
}
