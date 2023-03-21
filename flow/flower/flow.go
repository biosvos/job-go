package flower

import (
	"github.com/pkg/errors"
	"job-go/flow/recruiter"
	"job-go/flow/tagger"
)

var _ Flower = &Flow{}

type Flow struct {
	recruiter recruiter.Recruiter
	tagger    tagger.Tagger
}

func NewFlow(recruiter recruiter.Recruiter, tagger tagger.Tagger) *Flow {
	return &Flow{recruiter: recruiter, tagger: tagger}
}

func (f *Flow) ListJobs() ([]*Job, error) {
	jobs, err := f.recruiter.ListJobs(recruiter.WithMinAnnualIncome(7000))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []*Job
	for _, job := range jobs {
		tags := f.tagger.Tagging(job.Title)
		tags = append(tags, f.tagger.Tagging(job.Requirement)...)

		ret = append(ret, &Job{
			Title:       job.Title,
			Requirement: job.Requirement,
			Url:         job.Url,
			Tags:        tags,
		})
	}
	return ret, nil
}
