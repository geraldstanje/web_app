package main

import (
  "io"
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

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    homeHandler(w, r, b.Print())
  })

  http.ListenAndServe(":8080", nil)
}