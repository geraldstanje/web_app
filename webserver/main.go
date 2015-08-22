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
            <input type="file" id="file" name="file">
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