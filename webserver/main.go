package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "log"
)

func main() {
  fmt.Println("[database] Connecting to database...")
  _, err := sql.Open("postgres", "postgres://admin:changeme@192.168.59.103/admin?sslmode=verify-full") //?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("[database] Connected successfully.")
}