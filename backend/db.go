package backend

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "eventful.db")
	if err != nil {
		log.Fatal(err)
	}

	err = createTables()
	if err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	eventTable := `CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        date TEXT
    );`

	_, err := db.Exec(eventTable)
	return err
}
