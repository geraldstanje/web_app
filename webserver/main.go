package main

import (
  "io"
  "net/http"
  "log"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func homeHandler(w http.ResponseWriter, r *http.Request, msg string) {
  io.WriteString(w, msg)
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

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    homeHandler(w, r, "Hello World")
  })

  http.ListenAndServe(":8080", nil)
}