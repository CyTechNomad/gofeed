// go:generate fyne bundle -o bundle.go assets

package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
)

type goFeedTheme struct {
	fyne.Theme
}

func GetTheme() fyne.Theme {
	return &goFeedTheme{Theme: theme.DefaultTheme()}
}

func (g *goFeedTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
    return theme.DefaultTheme().Color(name, theme.VariantLight)
}   

func (g *goFeedTheme) Size(name fyne.ThemeSizeName) float32 {
    if name == theme.SizeNameText {
        return 12
    }

    return theme.DefaultTheme().Size(name)
}

