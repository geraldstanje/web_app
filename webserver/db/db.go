package db

import (
  "log"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func IsValidLogin(username string, password string) bool {
  fmt.Println("[database] Connecting to database...")
  db, err := sql.Open("postgres", "postgres://admin:changeme@192.168.59.103:5432/admin?sslmode=disable")
  defer db.Close()
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  } 
  fmt.Println("[database] Connected successfully.")

  query := fmt.Sprintf("SELECT password FROM account WHERE username='%s'", username)
  var pass string
  err = db.QueryRow(query).Scan(&pass)
  if err == sql.ErrNoRows {
    fmt.Println("[database] login failed...")
    return false
  } else if err != nil {
    log.Fatal(err)
  }

  if password == pass {
    return true
  }

  fmt.Println("[database] login failed...")
  return false
}