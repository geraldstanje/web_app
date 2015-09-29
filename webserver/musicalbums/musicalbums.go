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

type UserInfo struct {
	User string
}

type ImageInfo struct {
	Filename string
	Width    string
	Height   string
}

const imageLink = `<img border="5" style="margin:5px 5px" src="files/{{.Filename}}" width="{{.Width}}" height="{{.Height}}">`

func renderImgTemplate(w http.ResponseWriter, filename string, width string, height string) error {
	img := ImageInfo{Filename: filename, Width: width, Height: height}
	t, err := template.New("imagelink").Parse(imageLink)
	if err != nil {
		return err
	}
	err = t.Execute(w, img)
	return err
}

func renderMusicAlbumsTemplate(w http.ResponseWriter, user string) error {
	useraccount := UserInfo{User: user}
	t, err := template.ParseFiles("templates/musicalbums.html")
	if err != nil {
		return err
	}
	err = t.Execute(w, useraccount)
	return err
}

func formFile(r *Request) (multipart.File, error) {
  _, header, err := r.FormFile("TheFile")
  return header, err
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		header, error := formFile(r)
    if err != nil {
      log.Println("[webserver] " + err.Error())
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

		file, err := header.Open()
    if err != nil {
        log.Println("[webserver] " + err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
		path := fmt.Sprintf("files/%s", header.Filename)
		buf, err := ioutil.ReadAll(file)
    if err != nil {
        log.Println("[webserver] " + err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
		ioutil.WriteFile(path, buf, 0644)

		files, err := ioutil.ReadDir("./files")
    if err != nil {
      log.Println("[webserver] " + err.Error())
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

		for _, f := range files {
			err = renderImgTemplate(w, f.Name(), size, size)
			if err != nil {
				log.Println("[webserver] " + err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}
		}
	}
}

func Resize(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		size = r.FormValue("value")
		if size == "" {
			log.Println("[webserver] " + "Empty FormValues")
			http.Error(w, "Empty FormValue", http.StatusInternalServerError)
      return
		}

		files, err := ioutil.ReadDir("./files")
    if err != nil {
      log.Println("[webserver] " + err.Error())
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

		for _, f := range files {
			err := renderImgTemplate(w, f.Name(), size, size)
			if err != nil {
        log.Println("[webserver] " + err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
        return
			}
		}
	}
}

func MusicAlbums(w http.ResponseWriter, req *http.Request) {
	user, err := s.GetSessionUser(req)
	if user == "" || err != nil {
		http.Redirect(w, req, "/", 302)
		return
	}

	err = renderMusicAlbumsTemplate(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}