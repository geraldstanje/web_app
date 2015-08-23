package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  a "github.com/geraldstanje/web_app/webserver/authentication"
  m "github.com/geraldstanje/web_app/webserver/musicalbums"
)

var router = mux.NewRouter()

func indexHandler(response http.ResponseWriter, request *http.Request) {
  response.Header().Set("Content-type", "text/html")
  webpage, err := ioutil.ReadFile("templates/index.html")

  if err != nil {
    http.Error(response, fmt.Sprintf("index.html file error %v", err), 500)
  }

  fmt.Fprint(response, string(webpage))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  a.Login(w,r)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
  a.Logout(w,r)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  m.Upload(w,r)
}

func musicAlbumsHandler(w http.ResponseWriter, r *http.Request) {
  m.MusicAlbums(w,r)
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
  m.Resize(w,r)
}

func main() {
  http.Handle("/", router)
  http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files/"))))
  router.HandleFunc("/", indexHandler)
  router.HandleFunc("/musicalbums", musicAlbumsHandler)
  router.HandleFunc("/upload", uploadHandler)
  router.HandleFunc("/resize", resizeHandler)
  router.HandleFunc("/login", LoginHandler).Methods("POST")
  router.HandleFunc("/logout", LogoutHandler).Methods("POST")
  http.ListenAndServe(":8080", nil)
}