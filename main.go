package main

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Contact Data")

	db, err := sql.Open("sqlite3", "~/Desktop/raul-logistica/data-logistic.sqlite")
	if err != nil {
		panic(err)
	}
	var loadedData Agenda

	var list *widget.List

	list = widget.NewList(
		func() int { return len(loadedData) },
		func() fyne.CanvasObject {
			nombreLabel := widget.NewLabel("Nombre Empresa")
			telefLabel := widget.NewLabel("Telefono")
			zonaLabel := widget.NewLabel("Zona")
			contactContainer := container.NewGridWithColumns(4, nombreLabel, telefLabel, zonaLabel)
			edtb := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			delb := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			buttonContainer := container.NewHBox(edtb, delb)
			return container.NewBorder(nil, nil, nil, buttonContainer, contactContainer)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := o.(*fyne.Container)

			contactContainer := c.Objects[0].(*fyne.Container)
			buttonContainer := c.Objects[1].(*fyne.Container)

			nombreLabel := contactContainer.Objects[0].(*widget.Label)
			telefLabel := contactContainer.Objects[1].(*widget.Label)
			zonaLabel := contactContainer.Objects[2].(*widget.Label)

			edtb := buttonContainer.Objects[0].(*widget.Button)
			delb := buttonContainer.Objects[1].(*widget.Button)

			_ = loadedData[i].IdEntrada

			nombreLabel.SetText(loadedData[i].NombreEmpresa)
			telefLabel.SetText(loadedData[i].Telefono1)
			zonaLabel.SetText(loadedData[i].Zona)

			edtb.OnTapped = func() {

				entryFirstName := widget.NewEntry()
				entryZona := widget.NewEntry()
				entryPhone := widget.NewEntry()

				firstNameForm := widget.NewFormItem("Nombre Empresa", entryFirstName)
				lastNameForm := widget.NewFormItem("Zona", entryZona)
				phoneForm := widget.NewFormItem("Telefono", entryPhone)

				formItems := []*widget.FormItem{firstNameForm, lastNameForm, phoneForm}

				dialog1 := dialog.NewForm("Edit Contact", "Save", "Cancel", formItems, func(b bool) {

					///UPDATE THE DB WITH THE INPUT
					list.Refresh()
				}, myWindow)

				entryFirstName.SetText(nombreLabel.Text)
				entryZona.SetText(zonaLabel.Text)
				entryPhone.SetText(telefLabel.Text)

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
