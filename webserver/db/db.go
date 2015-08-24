package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
  "strings"
)

var dbIp = "192.168.59.103"
var dbName = "admin"
var dbUser = "admin"
var dbPass = "changeme"

func IsValidRegistration(user string, password string) bool {
	log.Println("[database] Connecting to database...")
	db, err := sql.Open("postgres", "postgres://" + dbUser + ":" + dbPass + "@" + dbIp + ":5432/" + dbName + "?sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[database] Connected successfully.")

	_, err = db.Query("INSERT INTO account VALUES($1, $2)", user, password)
  if err != nil {
    if strings.Contains(err.Error(), "violates unique contraint") {
      log.Println("[database] registeration failed, duplicated key")
		  return false
    } else {
      log.Fatal(err)
    }
	}

	return true
}

func IsValidLogin(user string, password string) bool {
	log.Println("[database] Connecting to database...")
	db, err := sql.Open("postgres", "postgres://" + dbUser + ":" + dbPass + "@" + dbIp + ":5432/" + dbName + "?sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[database] Connected successfully.")

	var pass string
	err = db.QueryRow("SELECT password FROM account WHERE email = $1", user).Scan(&pass)
	if err == sql.ErrNoRows {
		log.Println("[database] login failed...")
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	if password == pass {
		return true
	}

	log.Println("[database] login failed...")
	return false
}
