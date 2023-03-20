package programmers

import (
	"github.com/stretchr/testify/require"
	"job-go/flow/recruiter"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	jobs, err := NewProgrammers().ListJobs(recruiter.WithMinAnnualIncome(7000))
	require.NoError(t, err)
	log.Println(jobs)
}

func TestGetJob(t *testing.T) {
	job, err := requestJob(17452)
	require.NoError(t, err)
	content, err := newJobContent(job)
	require.NoError(t, err)
	t.Log(content.JobPosition.Requirement)
}
