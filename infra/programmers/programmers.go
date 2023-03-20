package programmers

import (
	"github.com/pkg/errors"
	"job-go/flow/recruiter"
)

var _ recruiter.Recruiter = &Programmers{}

type Programmers struct {
	repo Repository
}

func NewProgrammers(repository Repository) *Programmers {
	return &Programmers{repo: repository}
}

func (p *Programmers) ListJobs(opts ...recruiter.ListJobOption) ([]*recruiter.Job, error) {
	options := recruiter.ApplyListJobOptions(opts)
	content, err := getPageListJobContent(1, options)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret, err := content.extractJobs()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	totalPage := content.getTotalPage()
	for i := 2; i <= totalPage; i++ {
		content, err := getPageListJobContent(1, options)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		jobs, err := content.extractJobs()
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
	ret, err := newListJobsContent(body)
	return ret, errors.WithStack(err)
}
