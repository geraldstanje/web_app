package main

import (
  "io"
  "net/http"
  "bytes"
  "database/sql"
  _ "github.com/lib/pq"
)

type Logger struct {
  writer bytes.Buffer
}

func NewLogger() *Logger {
  return &Logger{}
}

func (l *Logger) EmitLine(line string) {
  l.writer.WriteString(line)
  l.writer.WriteString("\n")
}

func (l *Logger) Print() string {
  return l.writer.String()
}

func homeHandler(w http.ResponseWriter, r *http.Request, msg string) {
  io.WriteString(w, msg)
}

func main() {
  l := NewLogger()

  l.EmitLine("[database] Connecting to database...")
  db, err := sql.Open("postgres", "postgres://admin:changeme@192.168.59.103:5432/admin?sslmode=disable") //?sslmode=verify-full")
  if err != nil {
    l.EmitLine(err.Error())
  }
  err = db.Ping()
  if err != nil {
    l.EmitLine(err.Error())
  } 
  l.EmitLine("[database] Connected successfully.")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    homeHandler(w, r, l.Print())
  })

  http.ListenAndServe(":8000", nil)
}