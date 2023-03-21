package flower

type Job struct {
	Title       string
	Requirement string
	Url         string
	Tags        []string
}

type Flower interface {
	ListJobs() ([]*Job, error)
}
