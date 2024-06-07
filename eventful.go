package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"eventful/db"
	"log"
	"net/http"

	"eventful/handler"

	"github.com/rs/cors"
)

func Run() {
	db.InitDB()
	defer db.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/events", handler.GetEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", handler.GetEvent).Methods("GET")
	r.HandleFunc("/api/events", handler.CreateEvent).Methods("POST")
	r.HandleFunc("/api/events/{id}", handler.UpdateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{id}", handler.DeleteEvent).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	srv := &http.Server{
		Handler: c.Handler(r),
		Addr:    ":8080",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Server is running on port 8080")

	// Graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server shutting down")

}

func main() {
	Run()
}

// Remove the unused main function
