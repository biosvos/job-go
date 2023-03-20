package programmers

import (
	"encoding/json"
	"github.com/pkg/errors"
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

func newListJobsContent(body []byte) (*listJobsContent, error) {
	var ret listJobsContent
	err := json.Unmarshal(body, &ret)
	return &ret, errors.WithStack(err)
}

func (l *listJobsContent) extractJobs() ([]*recruiter.Job, error) {
	var ret []*recruiter.Job
	for _, position := range l.JobPositions {
		ret = append(ret, &recruiter.Job{
			Title: position.Title,
			Url:   "https://career.programmers.co.kr" + position.Url,
			Company: recruiter.Company{
				Name:    position.Company.Name,
				Address: position.Company.Address,
			},
		})
	}
	return ret, nil
}

func (l *listJobsContent) getTotalPage() int {
	return l.TotalPages
}

type jobContent struct {
	JobPosition struct {
		Id             int         `json:"id"`
		Address        string      `json:"address"`
		Career         string      `json:"career"`
		CareerRange    interface{} `json:"careerRange"`
		CompanyId      int         `json:"companyId"`
		JobType        string      `json:"jobType"`
		MaxSalary      int         `json:"maxSalary"`
		MinSalary      int         `json:"minSalary"`
		Personalized   interface{} `json:"personalized"`
		SigningBonus   int         `json:"signingBonus"`
		Status         string      `json:"status"`
		Title          string      `json:"title"`
		StartAt        time.Time   `json:"startAt"`
		CreatedAt      time.Time   `json:"createdAt"`
		UpdatedAt      time.Time   `json:"updatedAt"`
		CareerOption   bool        `json:"careerOption"`
		JobCategoryIds []int       `json:"jobCategoryIds"`
		Period         string      `json:"period"`
		TechnicalTags  []struct {
			Id            int     `json:"id"`
			Name          string  `json:"name"`
			TaggingsCount int     `json:"taggings_count"`
			Category      string  `json:"category"`
			Approved      bool    `json:"approved"`
			DisplayOrder  float64 `json:"display_order"`
		} `json:"technicalTags"`
		TeamTechnicalTags []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"teamTechnicalTags"`
		Url      string `json:"url"`
		IsSynced bool   `json:"isSynced"`
		Company  struct {
			Id                  int      `json:"id"`
			Name                string   `json:"name"`
			Address             string   `json:"address"`
			ServiceName         string   `json:"serviceName"`
			ServiceUrl          string   `json:"serviceUrl"`
			HomeUrl             string   `json:"homeUrl"`
			Funding             float64  `json:"funding"`
			HideFunding         bool     `json:"hideFunding"`
			Revenue             float64  `json:"revenue"`
			HideRevenue         bool     `json:"hideRevenue"`
			Blog                string   `json:"blog"`
			CountryCode         string   `json:"countryCode"`
			AverageResponseTime string   `json:"averageResponseTime"`
			BenefitTags         []string `json:"benefitTags"`
			EmployeesCount      int      `json:"employeesCount"`
			LogoUrl             string   `json:"logoUrl"`
			Developers          []struct {
				Url         string `json:"url"`
				Icon        string `json:"icon"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"developers"`
			Employees []struct {
				EmployeesCount int       `json:"employeesCount"`
				RegisteredAt   time.Time `json:"registeredAt"`
			} `json:"employees"`
		} `json:"company"`
		CountryCode           string      `json:"countryCode"`
		Latitude              float64     `json:"latitude"`
		Longitude             float64     `json:"longitude"`
		MinCareerRequired     bool        `json:"minCareerRequired"`
		MinCareer             interface{} `json:"minCareer"`
		ResumeRequired        bool        `json:"resumeRequired"`
		VideoUrl              interface{} `json:"videoUrl"`
		EndAt                 interface{} `json:"endAt"`
		AdditionalInformation string      `json:"additionalInformation"`
		Description           string      `json:"description"`
		PreferredExperience   string      `json:"preferredExperience"`
		Requirement           string      `json:"requirement"`
		TeamEnvironment       struct {
			DeveloperCount    string `json:"developerCount"`
			Os                string `json:"os"`
			VersionControl    string `json:"versionControl"`
			CollaborationTool string `json:"collaborationTool"`
			ReviewTool        string `json:"reviewTool"`
			ReviewMethod      string `json:"reviewMethod"`
		} `json:"teamEnvironment"`
		IsAppliable            bool `json:"isAppliable"`
		JobPositionCertificate struct {
			Status    interface{} `json:"status"`
			ExamType  interface{} `json:"examType"`
			Languages interface{} `json:"languages"`
			Level     interface{} `json:"level"`
		} `json:"jobPositionCertificate"`
	} `json:"jobPosition"`
	Settings struct {
		Google struct {
			JobPosition struct {
				Context            string    `json:"@context"`
				Type               string    `json:"@type"`
				DatePosted         time.Time `json:"datePosted"`
				Description        string    `json:"description"`
				EmploymentType     string    `json:"employmentType"`
				Title              string    `json:"title"`
				HiringOrganization struct {
					Type   string `json:"@type"`
					Name   string `json:"name"`
					SameAs string `json:"sameAs"`
					Logo   string `json:"logo"`
				} `json:"hiringOrganization"`
				JobLocation struct {
					Type    string `json:"@type"`
					Address string `json:"address"`
				} `json:"jobLocation"`
			} `json:"jobPosition"`
			TagManager struct {
				JobPositionId          int    `json:"jobPositionId"`
				JobPositionCategoryIds string `json:"jobPositionCategoryIds"`
			} `json:"tagManager"`
		} `json:"google"`
	} `json:"settings"`
}

func newJobContent(body []byte) (*jobContent, error) {
	var ret jobContent
	err := json.Unmarshal(body, &ret)
	return &ret, errors.WithStack(err)
}

func (j *jobContent) extractJob() *recruiter.Job {
	return &recruiter.Job{
		Title:       j.JobPosition.Title,
		Requirement: j.JobPosition.Requirement,
		Url:         "https://career.programmers.co.kr" + j.JobPosition.Url,
		Company: recruiter.Company{
			Name:    j.JobPosition.Company.Name,
			Address: j.JobPosition.Company.Address,
		},
	}
}
