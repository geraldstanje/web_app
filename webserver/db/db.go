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

  rows, err := db.Query("SELECT password FROM account WHERE username=?", username)
  if err != nil {
    fmt.Println("[database] login failed...")
    return false
  }

  defer rows.Close()
  for rows.Next() {
    var email string
    var user string
    var pass string

    err = rows.Scan(&pass)
    if err != nil {
      log.Fatal(err)
    }

    if password == pass {
      return true
    }
  }

  fmt.Println("[database] login failed...")
  return false
}