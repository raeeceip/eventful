package backend

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./eventful.db")
	if err != nil {
		log.Fatal(err)
	}

	err = createTables()
	if err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	userTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        role TEXT NOT NULL
    );`

	eventTable := `CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        date TEXT
    );`

	_, err := db.Exec(userTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(eventTable)
	if err != nil {
		return err
	}

	return nil
}
