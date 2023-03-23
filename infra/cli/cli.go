package cli

import (
	"fmt"
	"job-go/flow/flower"
)

type Cli struct {
	flower flower.Flower
}

func NewCli(flower flower.Flower) *Cli {
	return &Cli{flower: flower}
}

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

func (c *Cli) Run() {
	jobs, err := c.flower.ListJobs()
	if err != nil {
		panic(err)
	}

	for _, job := range jobs {
		if len(job.Tags) == 0 {
			fmt.Print(colorRed)
			fmt.Println(job.Title, job.Tags, job.Url)
			fmt.Print(colorReset)
		} else {
			fmt.Println(job.Title, job.Tags, job.Url)
		}
	}
}
