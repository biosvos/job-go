package recruiter

type Job struct {
	Title        string
	Requirements []string
	Url          string
	Company      Company
}

type Company struct {
	Name    string
	Address string
}
