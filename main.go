package main

import (
	"eventful/backend"
	"eventful/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Eventful")

	backend.InitAPI()

	w.Resize(fyne.NewSize(800, 600))
	ui.ShowSplashScreen(w)
	w.ShowAndRun()
}
