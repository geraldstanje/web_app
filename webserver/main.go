/*
package main

import (
  "io"
  "io/ioutil"
  "net/http"
  "log"
  "fmt"
  "bytes"
  "database/sql"
  _ "github.com/lib/pq"
)

func homeHandler(w http.ResponseWriter, r *http.Request, msg string) {
  io.WriteString(w, msg)
}

type Buffer struct {
  writer bytes.Buffer
}

func NewBuffer() *Buffer {
  return &Buffer{}
}

func (b *Buffer) EmitLine(line string) {
  b.writer.WriteString(line)
  b.writer.WriteString("\n")
}

func (b *Buffer) Print() string {
  return b.writer.String()
}

func upload(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    _, header, _ := r.FormFile("file")
    file, _ := header.Open()
    path := fmt.Sprintf("files/%s", header.Filename)
    buf, _ := ioutil.ReadAll(file)
    ioutil.WriteFile(path, buf, 0644)
    http.Redirect(w, r, "/"+path, 301)
  } else {
    http.Redirect(w, r, "/", 301)
  }
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `<html>
    <head>
        <title>Music album collection</title>
    </head>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <input type="file" id="file" name="file" accept="image/*">
            <input type="submit" name="submit" value="submit">
        </form>
    </body>
</html>`)
}

func main() {
  fmt.Println("[database] Connecting to database...")
  db, err := sql.Open("postgres", "postgres://admin:changeme@192.168.59.103:5432/admin?sslmode=disable") //?sslmode=verify-full")
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  } 
  fmt.Println("[database] Connected successfully.")

  rows, err := db.Query("SELECT * FROM account")
  if err != nil {
    log.Fatal(err)
  }

  b := NewBuffer()

  for rows.Next() {
    var email string
    var username string
    var password string

    err = rows.Scan(&email, &username, &password)
    if err != nil {
      log.Fatal(err)
    }
    b.EmitLine(email + " " + username + " " + password)
  }

  staticServer := http.StripPrefix("/files/", http.FileServer(http.Dir("files/")))
  http.HandleFunc("/", index)
  http.HandleFunc("/upload", upload)
  http.Handle("/files/", staticServer)

  //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //  homeHandler(w, r, b.Print())
  //})

  http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
  //"io"
  "io/ioutil"
  "net/http"
  "log"
  "fmt"
)

type MusicAlbumStore struct {
  size string
  errChan chan error
}

func NewMusicAlbumStore() *MusicAlbumStore {
  m := MusicAlbumStore{}
  m.size = "10"
  m.errChan = make(chan error) // unbuffered channel
  return &m
}

func (m *MusicAlbumStore) upload(w http.ResponseWriter, r *http.Request) {
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
      img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\"" + "files/%s" + "\" width=\"" + "%s" + "\" height=\"" + "%s" + "\">", f.Name(), m.size, m.size)
      w.Write([]byte(img))
    }
  }
 }

func (m *MusicAlbumStore) resize(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    m.size = r.FormValue("value")
    if m.size == "" {
      fmt.Println("Empty FormValue")
      return
    }

    files, _ := ioutil.ReadDir("./files")
    for _, f := range files {
      img := fmt.Sprintf("<img border=\"5\" style=\"margin:5px 5px\" src=\"" + "files/%s" + "\" width=\"" + "%s" + "\" height=\"" + "%s" + "\">", f.Name(), m.size, m.size)
      w.Write([]byte(img))
    }
  }
}

func (m *MusicAlbumStore) homeHandler(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  var text = `
<!DOCTYPE html>
<html>
<head>
<!--
<script src="https://code.jquery.com/jquery.min.js"></script>
<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
<script src="https://code.jquery.com/jquery-2.1.4.js"></script>
<script src="http://dn-xcdn-us.qbox.me/libs/bootstrap-filestyle/1.1.2/bootstrap-filestyle-min.js"></script>
-->
<script type="text/javascript">
var fileName = '';
function fileSelected() {
  try {
    var file = document.getElementById('TheFile').files[0];
    if (file) {
      fileName = file.name;
    }
  } catch(err) {
    //nothing
  }
  uploadFile();
}

function uploadFile() {
  try {
    var fd = new FormData();
    fd.append("TheFile", document.getElementById('TheFile').files[0]);
    var xhr = new XMLHttpRequest();
    //xhr.upload.addEventListener("progress", uploadProgress, false);
    //xhr.addEventListener("load", uploadComplete, false);
    //xhr.addEventListener("error", uploadFailed, false);
    //xhr.addEventListener("abort", uploadCanceled, false);
    xhr.open("POST", "/upload");
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        mycallback(xhr.responseText); // Another callback here
      }
    };
    xhr.send(fd);
  } catch(err) {
    document.getElementById("fileForm").submit();
  }
}

function mycallback(data) {
  //alert(data);
  //$('#result').html(data);
  document.getElementById("result").innerHTML = data;
}

function uploadProgress(event) {
  //if (evt.lengthComputable) {
  //  var percentComplete = Math.round(event.loaded * 100 / event.total);
  //  document.getElementById('progressNumber').innerHTML = percentComplete.toString() + '%';
  //}
}

function uploadComplete(event) {
  //document.getElementById('progressNumber').innerHTML = 'Upload Complete for ' + fileName;
}

function uploadFailed(event) {
  //document.getElementById('progressNumber').innerHTML = 'Error';
}

function uploadCanceled(event) {
  //document.getElementById('progressNumber').innerHTML = 'Upload canceled';
}

function resizeAlbumCover(data) {
  try {
    var fd = new FormData();
    fd.append("value", data);
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/resize");
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        mycallback(xhr.responseText); // Another callback here
      }
    };
    xhr.send(fd);
  } catch(err) {
    //
  }
}

window.onload = function() {
//jQuery(function($){
  resizeAlbumCover(fader.value);
}
</script>

<meta charset="utf-8">
<title>Music Album Collections</title>
</head>
<body>

<!--
<div class="row">
  <div class="col-sm-4" style="width: 50px; float: left;">
    <label for=fader>Size</label>
  </div>
  <div class="col-sm-4">
    <input type="range" min="50" max="300" value="10" id="fader" step="10" oninput="resizeAlbumCover(value)" style="width: 150px; float: left; margin: 5px;">
  </div>
  <div class="col-sm-4" style="width: 500px;">
    <input type="file" name="TheFile" id="TheFile" onchange="fileSelected()" style="width: 600px; height: 40px; background: white;"><BR>
    <input type="file" id="TheFile" class="filestyle" data-size="sm" onchange="fileSelected()"> 
  </div>
</div>
-->

<div class="row">
  <div class="col-sm-4" style="width: 50px; float: left;">
    <label for=fader>Size:</label>
  </div>
  <div class="col-sm-4" style="width: 200px; float: left;">
    <input type="range" min="50" max="300" value="10" id="fader" step="10" oninput="resizeAlbumCover(value)" style="width: 150px; float: left; margin: 5px;">
  </div>
  <div class="col-sm-4" style="width: 200px; float: left;">
    <label for=fader>New Music Album Upload:</label>
  </div>
  <div class="col-sm-4">
    <input type="file" name="TheFile" id="TheFile" onchange="fileSelected()" style="width: 600px; height: 40px; background: white;"><BR>
  </div>
</div>

<!-- <div id="progressNumber"></div> -->

<div id='result'></div>
<!--
<script>
  $(":file").filestyle({
    buttonText: "Album Cover"
  });
</script>
-->
</body>
</html>
  `
  w.Write([]byte(text))
}

func (m *MusicAlbumStore) startHTTPServer() {
  staticServer := http.StripPrefix("/files/", http.FileServer(http.Dir("files/")))
  http.HandleFunc("/", http.HandlerFunc(m.homeHandler))
  http.HandleFunc("/upload", m.upload)
  http.HandleFunc("/resize", m.resize)
  http.Handle("/files/", staticServer)
  err := http.ListenAndServe(":8080", nil)
  m.errChan <- err
}

func main() {
  m := NewMusicAlbumStore()
  m.startHTTPServer()
}
*/

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