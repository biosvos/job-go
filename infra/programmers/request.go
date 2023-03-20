package programmers

import (
	"github.com/biosvos/rest"
	"github.com/biosvos/rest/http"
	"github.com/pkg/errors"
)

func requestListJobs(page uint64, minAnnualIncome uint64) ([]byte, error) {
	ret, err := http.NewClient().Get("https://career.programmers.co.kr/api/job_positions", rest.WithQueries(map[string]string{
		"min_salary":         toString(minAnnualIncome),
		"order":              "recent",
		"page":               toString(page),
		"job_category_ids[]": "1",
	})).Execute()
	return ret, errors.WithStack(err)
}

func requestJob(jobId uint64) ([]byte, error) {
	ret, err := http.NewClient().Get("https://career.programmers.co.kr/api/job_positions/" + toString(jobId)).Execute()
	return ret, errors.WithStack(err)
}
