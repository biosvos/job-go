package main

import (
	"job-go/flow/flower"
	"job-go/infra/cli"
	"job-go/infra/gui"
	"job-go/infra/pipe"
	"job-go/infra/programmers"
	"job-go/infra/publish"
	"os"
)

func main() {
	_ = os.Mkdir("pro", 0700)
	repository := programmers.NewFileRepository("pro")
	recruiter := programmers.NewProgrammers(repository)
	tagger, err := pipe.NewPipeTagger("file.tag")
	if err != nil {
		panic(err)
	}
	flow := flower.NewFlow(recruiter, tagger)

	runtime := os.Getenv("RUNTIME")
	switch runtime {
	case "fyne":
		app := gui.NewGui(flow)
		app.Run()
	case "notion":
		notion := publish.NewNotion(flow)
		notion.Run()
	default:
		app := cli.NewCli(flow)
		app.Run()
	}
}
