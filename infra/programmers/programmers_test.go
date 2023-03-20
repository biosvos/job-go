//go:build programmers

package programmers

import (
	"github.com/stretchr/testify/require"
	"job-go/flow/recruiter"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	_ = os.Mkdir("pro", 0700)
	repository := NewFileRepository("pro")
	jobs, err := NewProgrammers(repository).ListJobs(recruiter.WithMinAnnualIncome(7000))
	require.NoError(t, err)
	for _, job := range jobs {
		t.Logf("%+v", job)
	}
}

func TestGetJob(t *testing.T) {
	job, err := requestJob(17452)
	require.NoError(t, err)
	content, err := newJobContent(job)
	require.NoError(t, err)
	t.Log(content.JobPosition.Requirement)
}
