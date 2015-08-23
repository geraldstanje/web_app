package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func IsValidLogin(username string, password string) bool {
	log.Println("[database] Connecting to database...")
	db, err := sql.Open("postgres", "postgres://admin:changeme@192.168.59.103:5432/admin?sslmode=disable")
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
	err = db.QueryRow("SELECT password FROM account WHERE username = $1", username).Scan(&pass)
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
