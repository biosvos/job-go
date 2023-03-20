package recruiter

type Job struct {
	Title       string
	Requirement string
	Url         string
	Company     Company
}

type Company struct {
	Name    string
	Address string
}
