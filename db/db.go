package db

import (
	"database/sql"
	"eventful/eventtypes"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
}

func CreateEvent(event *eventtypes.Event) error {
	_, err := db.Exec("INSERT INTO events (name, location, date) VALUES ($1, $2, $3)", event.Name, event.Location, event.Date)
	if err != nil {
		return err
	}

	return nil
}

func GetEvents() ([]eventtypes.Event, error) {
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []eventtypes.Event{}
	for rows.Next() {
		var event eventtypes.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.Date)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int) (*eventtypes.Event, error) {
	var event eventtypes.Event
	err := db.QueryRow("SELECT * FROM events WHERE id = $1", id).Scan(&event.ID, &event.Name, &event.Location, &event.Date)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(event *eventtypes.Event) error {
	_, err := db.Exec("UPDATE events SET name = $1, location = $2, date = $3 WHERE id = $4", event.Name, event.Location, event.Date, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int) error {
	_, err := db.Exec("DELETE FROM events WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}
