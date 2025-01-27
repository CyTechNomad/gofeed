package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"Sub1", "Sub2", "Sub3"}
var content = map[string]string{
    data[0]: "Content1",
    data[1]: "Content2",
    data[2]: "Content3",
}

var currentContent = widget.NewRichText()

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {}),
	)
	toolbar.Prepend(widget.NewToolbarSpacer())

	return container.NewStack(toolbar)
}

func makeSubcriptionsList() fyne.CanvasObject {
	return widget.NewList(func() int { return len(data) }, func() fyne.CanvasObject { return widget.NewButton("template",func(){}) }, func(i int, o fyne.CanvasObject) { o.(*widget.Button).SetText(GetContent(data[i])) })
}

func getContent() fyne.CanvasObject {
    return currentContent
}

func makeGUI() fyne.CanvasObject {
	return container.NewBorder(makeBanner(), nil, makeSubcriptionsList(), nil, getContent())
}
