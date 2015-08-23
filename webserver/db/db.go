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

  rows, err := db.Query("SELECT * FROM account")
  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()
  for rows.Next() {
    var email string
    var user string
    var pass string

    fmt.Println("[database] check rows...")

    err = rows.Scan(&email, &user, &pass)
    if err != nil {
      log.Fatal(err)
    }
    
    fmt.Sprintf("[database] username: %d", len(username))
    fmt.Sprintf("[database] user: %d", len(user))
    fmt.Sprintf("[database] password: %d", len(password))
    fmt.Sprintf("[database] pass: %d", len(pass))

    if (username == user) && (password == pass) {
      fmt.Println("[database] login correct...")
      return true
    }
  }

  fmt.Println("[database] login failed...")
  return false
}