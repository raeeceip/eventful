package main

import (
	"context"
	"embed"
	"eventful/backend"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting application...")

	err = wails.Run(&options.App{
		Title:  "Eventful",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Println("Application stopped.")
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	log.Println("Running startup...")
	a.ctx = ctx
	backend.InitDB()
	backend.InitAPI()
	log.Println("Startup completed.")
}

func (a *App) GetEvents() []backend.Event {
	return backend.GetEvents()
}

func (a *App) CreateEvent(event backend.Event) error {
	return backend.CreateEvent(event)
}
