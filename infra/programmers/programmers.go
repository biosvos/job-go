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
	totalPage := 1
	run := true
	page := totalPage
	for run && page <= totalPage {
		content, err := getPageListJobContent(uint64(page), options)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		totalPage = content.getTotalPage()
		page++

		for _, id := range content.getJobIds() {
			success, err := p.saveNewJob(id)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if !success { // 이미 존재한다.
				run = false
				break
			}
		}
	}
	ret, err := p.listAllJobs()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, errors.WithStack(err)
}

// saveNewJob return save success
func (p *Programmers) saveNewJob(id int) (bool, error) {
	if p.repo.IsExists(jobFileName(id)) {
		return false, nil
	}
	body, err := requestJob(uint64(id))
	if err != nil {
		return false, errors.WithStack(err)
	}
	err = p.repo.Save(jobFileName(id), body)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}

func (p *Programmers) listAllJobs() ([]*recruiter.Job, error) {
	list, err := p.repo.List()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret []*recruiter.Job
	for _, item := range list {
		bytes, err := p.repo.Load(item)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		content, err := newJobContent(bytes)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		job := content.extractJob()
		ret = append(ret, job)
	}
	return ret, nil
}

func jobFileName(id int) string {
	return toString(uint64(id)) + ".json"
}

func getPageListJobContent(page uint64, options *recruiter.ListJobOptions) (*listJobsContent, error) {
	body, err := requestListJobs(page, options.MinAnnualIncome)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret, err := newListJobsContent(body)
	return ret, errors.WithStack(err)
}
