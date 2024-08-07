package backend

import (
	"log"
	"math/rand/v2"
	"net/http"
	"sort"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var e *echo.Echo
var jwtSecret = []byte("your-secret-key")

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
	// events routes
	e.GET("/events", getEventsHandler, jwtMiddleware)
	e.POST("/events", createEventHandler, jwtMiddleware)
	e.GET("/events/recommended", getRecommendedEventsHandler, jwtMiddleware)

	//login
	e.POST("/login", loginHandler)
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
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	score       float64 `json:"-"` // Internal use only, not exposed in JSON
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.ErrUnauthorized
		}

		token = token[7:] // Remove "Bearer " prefix

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func loginHandler(c echo.Context) error {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		return err
	}

	// Here you should check the credentials against your database
	// For this example, we'll use a hardcoded user
	if credentials.Username != "admin" || credentials.Password != "password" {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"token": t,
	})
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
func CreateEvent(event Event) error {
	_, err := db.Exec("INSERT INTO events (title, description, date) VALUES (?, ?, ?)",
		event.Title, event.Description, event.Date)
	return err
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

func getEventsHandler(c echo.Context) error {
	events := GetEvents()
	return c.JSON(200, events)
}

func createEventHandler(c echo.Context) error {
	var event Event
	if err := c.Bind(&event); err != nil {
		return err
	}
	err := CreateEvent(event)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	return c.JSON(201, event)
}

func getRecommendedEventsHandler(c echo.Context) error {
	allEvents := GetEvents()
	recommendedEvents := recommendEvents(allEvents)
	return c.JSON(200, recommendedEvents)
}

func recommendEvents(events []Event) []Event {
	// Simple recommendation algorithm:
	// 1. Sort events by date (most recent first)
	// 2. Calculate a score for each event based on recency and a random factor
	// 3. Sort events by score and return top 5

	sort.Slice(events, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", events[i].Date)
		dateJ, _ := time.Parse("2006-01-02", events[j].Date)
		return dateI.After(dateJ)
	})

	now := time.Now()
	for i := range events {
		date, _ := time.Parse("2006-01-02", events[i].Date)
		daysDiff := now.Sub(date).Hours() / 24
		randomFactor := rand.Float64() // Add some randomness to recommendations
		events[i].score = 1/(daysDiff+1) + randomFactor
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].score > events[j].score
	})

	numRecommended := 5
	if len(events) < numRecommended {
		numRecommended = len(events)
	}
	return events[:numRecommended]
}
