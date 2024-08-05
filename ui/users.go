package ui

import (
	"eventful/backend"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UsersScreen(w fyne.Window) fyne.CanvasObject {
	users := backend.GetUsers()

	list := widget.NewList(
		func() int { return len(users) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template User")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			user := users[id]
			item.(*widget.Label).SetText(fmt.Sprintf("%s (%s)", user.Username, user.Role))
		},
	)

	addButton := widget.NewButton("Add User", func() {
		showAddUserDialog(w, list)
	})

	backButton := widget.NewButton("Back", func() {
		w.SetContent(HomeScreen(w))
	})

	return container.NewBorder(addButton, backButton, nil, nil, list)
}

func showAddUserDialog(w fyne.Window, list *widget.List) {
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	role := widget.NewSelect([]string{"user", "organizer", "admin"}, nil)
	role.SetSelected("user")

	dialog := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
		widget.NewFormItem("Role", role),
	)

	dialog.OnSubmit = func() {
		user := &backend.User{
			Username: username.Text,
			Password: password.Text,
			Role:     role.Selected,
		}
		backend.CreateUser(user)
		list.Refresh()
	}

	dialog.OnCancel = func() {
		// Do nothing, dialog will be dismissed
	}

	w.SetContent(dialog)
}
