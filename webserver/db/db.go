package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

const (
	DB_USER     = "admin"
	DB_PASSWORD = "changeme"
	DB_HOST     = "postgresql_db"
	DB_NAME     = "admin"
)

func AddUser(user string, password string) bool {
	log.Println("[database] Connecting to database...")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db, err := sql.Open("postgres", dbinfo)
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
		if strings.Contains(err.Error(), "violates unique constraint") {
			log.Println("[database] registeration failed, duplicated key...")
			return false
		} else {
			log.Fatal(err)
		}
	}

	return true
}

func CheckUserLogin(user string, password string) bool {
	log.Println("[database] Connecting to database...")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db, err := sql.Open("postgres", dbinfo)
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

func RemoveUser(user string) bool {
  log.Println("[database] Connecting to database...")

  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
  db, err := sql.Open("postgres", dbinfo)
  defer db.Close()
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("[database] Connected successfully.")

  res, err := db.Exec("DELETE FROM account WHERE email = $1", user)
  if err != nil {
    log.Fatal(err)
  } 

  changes, err := res.RowsAffected()
  if err != nil {
    log.Fatal(err)
  } 

  if changes == 0 {
    return false
  }

  return true
}