package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(GetTheme())

	w := a.NewWindow("Go Feed")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(makeGUI())
	w.ShowAndRun()
}
