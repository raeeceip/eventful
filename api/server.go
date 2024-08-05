package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

var events = []Event{
	{ID: "1", Title: "Conference", Description: "Tech conference", Date: "2024-09-15"},
	{ID: "2", Title: "Workshop", Description: "Coding workshop", Date: "2024-10-01"},
}

func StartServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${method} ${uri}\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/events", getEvents)
	e.GET("/events/:id", getEvent)
	e.POST("/events", createEvent)
	e.PUT("/events/:id", updateEvent)
	e.DELETE("/events/:id", deleteEvent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers
func getEvents(c echo.Context) error {
	return c.JSON(http.StatusOK, events)
}

func getEvent(c echo.Context) error {
	id := c.Param("id")
	for _, event := range events {
		if event.ID == id {
			return c.JSON(http.StatusOK, event)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Event not found"})
}

func createEvent(c echo.Context) error {
	event := new(Event)
	if err := c.Bind(event); err != nil {
		return err
	}
	events = append(events, *event)
	return c.JSON(http.StatusCreated, event)
}

func updateEvent(c echo.Context) error {
	id := c.Param("id")
	for i, event := range events {
		if event.ID == id {
			if err := c.Bind(&events[i]); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, events[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Event not found"})
}

func deleteEvent(c echo.Context) error {
	id := c.Param("id")
	for i, event := range events {
		if event.ID == id {
			events = append(events[:i], events[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Event not found"})
}
