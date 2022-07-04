package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Contact struct {
	ID        string `json:"_id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Contact Data")

	var loadedData []Contact

	var list *widget.List

	list = widget.NewList(
		func() int { return len(loadedData) },
		func() fyne.CanvasObject {
			firstNameLabel := widget.NewLabel("First Name")
			lastNameLabel := widget.NewLabel("Last Name")
			emailLabel := widget.NewLabel("Email")
			phoneLabel := widget.NewLabel("Phone")
			contactContainer := container.NewGridWithColumns(4, firstNameLabel, lastNameLabel, emailLabel, phoneLabel)
			edtb := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			delb := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			buttonContainer := container.NewHBox(edtb, delb)
			return container.NewBorder(nil, nil, nil, buttonContainer, contactContainer)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)

			contactContainer := c.Objects[0].(*fyne.Container)
			buttonContainer := c.Objects[1].(*fyne.Container)

			firstNameLabel := contactContainer.Objects[0].(*widget.Label)
			lastNameLabel := contactContainer.Objects[1].(*widget.Label)
			emailLabel := contactContainer.Objects[2].(*widget.Label)
			phoneLabel := contactContainer.Objects[3].(*widget.Label)

			edtb := buttonContainer.Objects[0].(*widget.Button)
			delb := buttonContainer.Objects[1].(*widget.Button)

			_ = loadedData[i].ID

			firstNameLabel.SetText(loadedData[i].FirstName)
			lastNameLabel.SetText(loadedData[i].LastName)
			emailLabel.SetText(loadedData[i].Email)
			phoneLabel.SetText(loadedData[i].Phone)

			edtb.OnTapped = func() {

				entryFirstName := widget.NewEntry()
				entryLastName := widget.NewEntry()
				entryEmail := widget.NewEntry()
				entryPhone := widget.NewEntry()

				firstNameForm := widget.NewFormItem("First Name", entryFirstName)
				lastNameForm := widget.NewFormItem("Last Name", entryLastName)
				emailForm := widget.NewFormItem("Email", entryEmail)
				phoneForm := widget.NewFormItem("Phone", entryPhone)

				formItems := []*widget.FormItem{firstNameForm, lastNameForm, emailForm, phoneForm}

				dialog1 := dialog.NewForm("Edit Contact", "Save", "Cancel", formItems, func(b bool) {

					///UPDATE THE DB WITH THE INPUT
					list.Refresh()
				}, myWindow)

				entryFirstName.SetText(firstNameLabel.Text)
				entryLastName.SetText(lastNameLabel.Text)
				entryEmail.SetText(emailLabel.Text)
				entryPhone.SetText(phoneLabel.Text)

				dialog1.Resize(fyne.NewSize(500, 300))

				dialog1.Show()

			}

			delb.OnTapped = func() {

				dialog1 := dialog.NewConfirm(
					"Delete Contact",
					"Do you wish to delete this contact?",
					func(b bool) {
						if b {
							//DELETE CONTACT FROM DB
						}
						list.Refresh()
					},
					myWindow,
				)

				dialog1.Resize(fyne.NewSize(300, 200))
				dialog1.Show()

			}

		})

	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
	}

	add := widget.NewButton("Add Contact", func() {

		entryFirstName := widget.NewEntry()
		entryLastName := widget.NewEntry()
		entryEmail := widget.NewEntry()
		entryPhone := widget.NewEntry()

		firstNameForm := widget.NewFormItem("First Name", entryFirstName)
		lastNameForm := widget.NewFormItem("Last Name", entryLastName)
		emailForm := widget.NewFormItem("Email", entryEmail)
		phoneForm := widget.NewFormItem("Phone", entryPhone)

		formItems := []*widget.FormItem{firstNameForm, lastNameForm, emailForm, phoneForm}

		dialogAdd := dialog.NewForm("Add Contact", "Add", "Cancel", formItems, func(b bool) {
			if b {
				_ = Contact{
					FirstName: entryFirstName.Text,
					LastName:  entryLastName.Text,
					Email:     entryEmail.Text,
					Phone:     entryPhone.Text,
				}

				//CREATE IN DB
				list.Refresh()
			}
		}, myWindow)

		dialogAdd.Resize(fyne.NewSize(500, 300))

		dialogAdd.Show()

	})

	exit := widget.NewButton("Quit", func() {

		myWindow.Close()
	})

	myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
