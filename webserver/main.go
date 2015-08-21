package main

import (
  "io"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

type struct Logger {
  msg string
}

func NewLogger() *Logger {
  return &Logger{}
}

func (l *Logger) EmitLine(line string) {
  l.writer.WriteString(line)
  l.writer.WriteString("\n")
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

  http.HandleFunc("/", makeHandler(homeHandler))
  http.ListenAndServe(":8000", nil)
}