// ui/events.go
package ui

import (
	"eventful/backend"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func EventsScreen(w fyne.Window) fyne.CanvasObject {
	events := backend.GetEvents()

	title := canvas.NewText("Events", theme.PrimaryColor())
	title.TextSize = 24
	title.Alignment = fyne.TextAlignCenter

	list := widget.NewList(
		func() int { return len(events) },
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Event"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(events[id].Title)
		},
	)

	backBtn := widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {
		w.SetContent(HomeScreen(w))
	})

	content := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()),
		backBtn,
		nil,
		nil,
		list,
	)

	return container.NewPadded(content)
}
