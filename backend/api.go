package backend

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var e *echo.Echo

func InitAPI() {
	InitDB()
	addSampleEvents()
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.GET("/users", getUsers)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server in a goroutine
	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}
type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func GetUsers() []User {
	rows, err := db.Query("SELECT id, username, role FROM users")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}

	return users
}

func CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)",
		user.Username, string(hashedPassword), user.Role)
	return err
}
func GetEvents() []Event {
	rows, err := db.Query("SELECT id, title, description, date FROM events")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Date); err != nil {
			log.Println(err)
			continue
		}
		events = append(events, e)
	}

	return events
}

func getEvents(c echo.Context) error {
	events := GetEvents()
	return c.JSON(http.StatusOK, events)
}
func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)",
		u.Username, string(hashedPassword), u.Role)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func getUsers(c echo.Context) error {
	rows, err := db.Query("SELECT id, username, role FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			return err
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	_, err := db.Exec("UPDATE users SET username = ?, role = ? WHERE id = ?",
		u.Username, u.Role, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
func addSampleEvents() {
	events := []Event{
		{Title: "Tech Conference", Description: "Annual technology conference", Date: "2024-09-15"},
		{Title: "Coding Workshop", Description: "Learn to code in Go", Date: "2024-10-01"},
	}

	for _, event := range events {
		_, err := db.Exec("INSERT INTO events (title, description, date) VALUES (?, ?, ?)",
			event.Title, event.Description, event.Date)
		if err != nil {
			log.Printf("Error adding sample event: %v", err)
		}
	}
}
