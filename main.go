package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func main() {
	var app fyne.App = app.New()
	var window fyne.Window = app.NewWindow("ADB GUI")

	var clickCounter int = 0

	var label *widget.Label = widget.NewLabel(fmt.Sprintf("You clicked the button %d times.", clickCounter))
	var button *widget.Button = widget.NewButton("Click here.", func() {
		clickCounter++
		label.SetText(fmt.Sprintf("You clicked the button %d times.", clickCounter))
	})

	var container *fyne.Container = container.New(
		layout.NewVBoxLayout(),
		button,
		// layout.NewSpacer(),
		label,
	)

	window.SetContent(container)
	window.ShowAndRun()
}