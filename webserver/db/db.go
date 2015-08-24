package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func IsValidRegistration(user string, password string) bool {
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

	_, err = db.Query("INSERT INTO account VALUES($1, $2)", user, password)
  log.Println("[database] error: " + err.Error())

	//if err == sql. {
	//	log.Println("[database] registeration failed...")
	//	return false
	//} else 
  if err != nil {
		return false
    //log.Fatal(err)
	}

	return true
}

func IsValidLogin(user string, password string) bool {
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
