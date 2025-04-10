package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/layout"
)

var (
    title = &widget.TextSegment{
        Style: widget.RichTextStyleHeading,
    }
    summary = &widget.TextSegment{
    }
    body = &widget.TextSegment{
    }
    contentWindow = widget.NewRichText(title, summary, body)
)

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), RefreshFeeds),
	)
	toolbar.Prepend(widget.NewToolbarSpacer())

	return container.NewStack(toolbar)
}

func makeSubcriptionsList() fyne.CanvasObject {
	subs := make([]string, 0, len(feeds))
	for sub := range feeds {
		subs = append(subs, sub)
	}

	list := widget.NewList(
		func() int { return len(feeds) },
		func() fyne.CanvasObject {
			return widget.NewButton("template", func() {})
		},
		func(i int, o fyne.CanvasObject) {
			button := o.(*widget.Button)
			button.SetText(feeds[subs[i]].Title)
			button.OnTapped = func() {
				if feed, ok := feeds[subs[i]]; ok && feed != nil {
                    title.Text = feed.Items[0].Title	
                    //summary.Text = feed.Items[0].Summary
                    if content := feed.Items[0].Content; content != "" {
                        body.Text = content
                    } else {
                        body.Text = scrapeTextContent(feed.Items[0].Link)
                    }
                        
                    contentWindow.Refresh()
				} else {
                    // This should never happen
                }
			}
		},
	)
	return list
}

func getContent() fyne.CanvasObject {
	return container.NewVScroll(container.New(layout.NewVBoxLayout(), contentWindow))
}

func makeGUI() fyne.CanvasObject {
	return container.NewBorder(makeBanner(), fyne.NewContainer(), makeSubcriptionsList(), nil, getContent())
}
