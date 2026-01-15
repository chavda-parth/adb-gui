package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ADB Test")

	myWindow.SetContent(widget.NewLabel("Hello. Testing ADB GUI"))

	myWindow.ShowAndRun()
}