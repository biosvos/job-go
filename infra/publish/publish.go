package publish

import (
	"context"
	"github.com/dstotijn/go-notion"
	"job-go/flow/flower"
	"os"
)

type Notion struct {
	flower flower.Flower
}

func NewNotion(flower flower.Flower) *Notion {
	return &Notion{flower: flower}
}

func (n *Notion) Run() {
	notionApiKey := os.Getenv("NOTION_API_KEY")
	notionDatabaseId := os.Getenv("NOTION_DATABASE_ID")
	client := notion.NewClient(notionApiKey)

	jobs, err := n.flower.ListJobs()
	if err != nil {
		panic(err)
	}
	for _, job := range jobs {
		var tags []notion.SelectOptions
		for _, tag := range job.Tags {
			tags = append(tags, notion.SelectOptions{
				Name: tag,
			})
		}
		if len(tags) > 0 {
			_, err := client.CreatePage(context.Background(), notion.CreatePageParams{
				ParentType: notion.ParentTypeDatabase,
				ParentID:   notionDatabaseId,
				DatabasePageProperties: &notion.DatabasePageProperties{
					"제목": {
						Title: []notion.RichText{
							{
								Text: &notion.Text{Content: job.Title},
							},
						},
					},
					"태그": {
						MultiSelect: tags,
					},
					"URL": {
						URL: &job.Url,
					},
				},
			})

			if err != nil {
				panic(err)
			}
		} else {
			_, err := client.CreatePage(context.Background(), notion.CreatePageParams{
				ParentType: notion.ParentTypeDatabase,
				ParentID:   notionDatabaseId,
				DatabasePageProperties: &notion.DatabasePageProperties{
					"제목": {
						Title: []notion.RichText{
							{
								Text: &notion.Text{Content: job.Title},
							},
						},
					},
					"URL": {
						URL: &job.Url,
					},
				},
			})
			if err != nil {
				panic(err)
			}
		}

	}
}
