package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type App struct {
	fyne.App
}

func NewApp() *App {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	return &App{
		App: a,
	}
}

var _ fyne.Theme = &myTheme{}

type myTheme struct{}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceNanumGothicTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceNanumGothicTtf
}
