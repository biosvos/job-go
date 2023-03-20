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
