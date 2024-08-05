package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen(w fyne.Window) fyne.CanvasObject {
	eventsButton := widget.NewButton("Events", func() {
		w.SetContent(EventsScreen(w))
	})

	usersButton := widget.NewButton("Users", func() {
		w.SetContent(UsersScreen(w))
	})

	content := container.NewVBox(
		widget.NewLabel("Welcome to Eventful"),
		eventsButton,
		usersButton,
	)

	return content
}
