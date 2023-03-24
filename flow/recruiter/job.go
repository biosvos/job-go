package recruiter

type Job struct {
	Title                     string
	QualificationRequirements string
	PreferredRequirements     string
	Url                       string
	Company                   Company
}

type Company struct {
	Name    string
	Address string
}
