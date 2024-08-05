// ui/splash.go
package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func ShowSplashScreen(w fyne.Window) {
	logo := canvas.NewText("Eventful", theme.PrimaryColor())
	logo.TextSize = 36
	logo.Alignment = fyne.TextAlignCenter

	splash := container.NewCenter(logo)
	w.SetContent(splash)
	w.Show()

	// Simulate loading time
	go func() {
		time.Sleep(2 * time.Second)
		w.SetContent(HomeScreen(w))
	}()
}
