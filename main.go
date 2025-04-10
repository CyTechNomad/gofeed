package main

import (
    "os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
    os.Setenv("FYNE_RENDERER", "software")
    loadFeeds(feeds)
	a := app.New()
	a.Settings().SetTheme(GetTheme())

	w := a.NewWindow("Go Feed")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(makeGUI())
	w.ShowAndRun()
}
