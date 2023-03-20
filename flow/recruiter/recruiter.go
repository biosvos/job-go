package recruiter

type ListJobOptions struct {
	MinAnnualIncome uint64
}

type ListJobOption func(options *ListJobOptions)

func WithMinAnnualIncome(income uint64) ListJobOption {
	return func(options *ListJobOptions) {
		options.MinAnnualIncome = income
	}
}

func ApplyListJobOptions(options []ListJobOption) *ListJobOptions {
	ret := ListJobOptions{
		MinAnnualIncome: 0,
	}
	for _, option := range options {
		option(&ret)
	}
	return &ret
}

type Recruiter interface {
	ListJobs(opts ...ListJobOption) ([]*Job, error)
}
