package programmers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"job-go/flow/recruiter"
)

var _ recruiter.Recruiter = &Programmers{}

type Programmers struct {
}

func NewProgrammers() *Programmers {
	return &Programmers{}
}

func (p *Programmers) ListJobs(opts ...recruiter.ListJobOption) ([]*recruiter.Job, error) {
	options := recruiter.ApplyListJobOptions(opts)
	content, err := getPageListJobContent(1, options)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret, err := extractJobs(content)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for i := 2; i <= content.TotalPages; i++ {
		content, err := getPageListJobContent(1, options)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		jobs, err := extractJobs(content)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, jobs...)
	}

	return ret, nil
}

func getPageListJobContent(page uint64, options *recruiter.ListJobOptions) (*listJobsContent, error) {
	body, err := requestListJobs(page, options.MinAnnualIncome)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var content listJobsContent
	err = json.Unmarshal(body, &content)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &content, nil
}
