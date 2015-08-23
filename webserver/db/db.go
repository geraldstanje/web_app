package db

import (
  "log"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func IsCorrectLogin(username string, password string) {
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

  for rows.Next() {
    var email string
    var user string
    var password string

    err = rows.Scan(&email, &user, &pass)
    if err != nil {
      log.Fatal(err)
    }
    
    if username == user && password == pass {
      return true
    }
  }

  return false
}