package flower

type Job struct {
	Title                     string
	QualificationRequirements string
	PreferredRequirements     string
	Url                       string
	Tags                      []string
}

type Flower interface {
	ListJobs() ([]*Job, error)
}
