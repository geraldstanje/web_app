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

const musicAlbumsPage = `
<!DOCTYPE html>
<html>
<head>
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

<small>User: {{.userName}}</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form><BR>

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
</body>
</html>
`
type User struct {
  userName string
}

func MusicAlbums(w http.ResponseWriter, req *http.Request) {
  userName := s.GetUserName(req)
  //fmt.Println("userName: " + userName)
  if userName != "" {
    //t := template.New("some template") // Create a template.
    //t, _ = t.ParseFiles("tmpl/welcome.html", nil)  // Parse template file.
    //user := GetUser() // Get current user infomration.

    user := User{userName: userName}

    t, _ := template.New("foo").Parse(musicAlbumsPage)
    t.Execute(w, user)  // merge.
    //fmt.Fprintf(w, musicAlbumsPage, userName)
  } else {
    http.Redirect(w, req, "/", 302)
  }
}