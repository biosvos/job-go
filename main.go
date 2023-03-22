package main

import (
	"job-go/flow/flower"
	"job-go/infra/gui"
	"job-go/infra/pipe"
	"job-go/infra/programmers"
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
	app := gui.NewGui(flow)
	app.Run()
}
