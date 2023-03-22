package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"job-go/flow/flower"
	"os/exec"
	"strings"
)

type Gui struct {
	app    *App
	flower flower.Flower
}

func NewGui(flower flower.Flower) *Gui {
	return &Gui{
		app:    NewApp(),
		flower: flower,
	}
}

func (g *Gui) Run() {
	jobs, err := g.flower.ListJobs()
	if err != nil {
		panic(err)
	}

	var cards []fyne.CanvasObject
	for _, job := range jobs {
		card := widget.NewCard(job.Title, strings.Join(job.Tags, ", "), widget.NewButton(job.Url, func() {
			_ = exec.Command("xdg-open", job.Url).Start()
		}))
		cards = append(cards, container.NewMax(card))
	}

	flow := container.NewGridWrap(fyne.NewSize(800, 200), cards...)
	scroll := container.NewVScroll(flow)
	window := g.app.NewWindow("search route")
	window.SetContent(scroll)
	window.ShowAndRun()
}
