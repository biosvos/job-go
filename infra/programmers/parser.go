package programmers

import (
	"job-go/flow/recruiter"
	"time"
)

type listJobsContent struct {
	JobPositions []struct {
		Id             int       `json:"id"`
		Address        string    `json:"address"`
		Career         string    `json:"career"`
		CareerRange    *string   `json:"careerRange"`
		CompanyId      int       `json:"companyId"`
		JobType        string    `json:"jobType"`
		MaxSalary      int       `json:"maxSalary"`
		MinSalary      int       `json:"minSalary"`
		SigningBonus   int       `json:"signingBonus"`
		Status         string    `json:"status"`
		Title          string    `json:"title"`
		StartAt        time.Time `json:"startAt"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		CareerOption   bool      `json:"careerOption"`
		JobCategoryIds []int     `json:"jobCategoryIds"`
		Period         string    `json:"period"`
		TechnicalTags  []struct {
			Id            int     `json:"id"`
			Name          string  `json:"name"`
			TaggingsCount int     `json:"taggings_count"`
			Category      string  `json:"category"`
			Approved      bool    `json:"approved"`
			DisplayOrder  float64 `json:"display_order"`
		} `json:"technicalTags"`
		TeamTechnicalTags []string `json:"teamTechnicalTags"`
		Url               string   `json:"url"`
		IsSynced          bool     `json:"isSynced"`
		Company           struct {
			Id                  int      `json:"id"`
			Name                string   `json:"name"`
			Address             string   `json:"address"`
			ServiceName         *string  `json:"serviceName"`
			ServiceUrl          *string  `json:"serviceUrl"`
			HomeUrl             string   `json:"homeUrl"`
			Funding             *float64 `json:"funding"`
			HideFunding         bool     `json:"hideFunding"`
			Revenue             *float64 `json:"revenue"`
			HideRevenue         bool     `json:"hideRevenue"`
			Blog                *string  `json:"blog"`
			CountryCode         string   `json:"countryCode"`
			AverageResponseTime *string  `json:"averageResponseTime"`
			EmployeesCount      int      `json:"employeesCount"`
			LogoUrl             string   `json:"logoUrl"`
			Developers          []struct {
				Url         string `json:"url"`
				Icon        string `json:"icon"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"developers"`
		} `json:"company"`
	} `json:"jobPositions"`
	FilterParams struct {
		Order          string `json:"order"`
		MinSalary      string `json:"min_salary"`
		JobCategoryIds []int  `json:"job_category_ids"`
	} `json:"filterParams"`
	Page         string `json:"page"`
	TotalPages   int    `json:"totalPages"`
	TotalEntries int    `json:"totalEntries"`
}

func extractJobs(content *listJobsContent) ([]*recruiter.Job, error) {
	var ret []*recruiter.Job
	for _, position := range content.JobPositions {
		ret = append(ret, &recruiter.Job{
			Title:        position.Title,
			Requirements: nil,
			Url:          "https://career.programmers.co.kr" + position.Url,
			Company: recruiter.Company{
				Name:    position.Company.Name,
				Address: position.Company.Address,
			},
		})
	}
	return ret, nil
}
