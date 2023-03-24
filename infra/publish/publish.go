package publish

import (
	"context"
	"github.com/dstotijn/go-notion"
	"job-go/flow/flower"
	"log"
	"os"
	"time"
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
		properties := notion.DatabasePageProperties{
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
		}

		if len(tags) > 0 {
			properties["태그"] = notion.DatabasePageProperty{
				MultiSelect: tags,
			}
		}

		param := notion.CreatePageParams{
			ParentType:             notion.ParentTypeDatabase,
			ParentID:               notionDatabaseId,
			DatabasePageProperties: &properties,
			Children: []notion.Block{
				notion.Heading1Block{
					RichText: []notion.RichText{{Text: &notion.Text{Content: "자격 조건"}}},
				},
				notion.ParagraphBlock{
					RichText: []notion.RichText{{Text: &notion.Text{Content: job.QualificationRequirements}}},
				},
				notion.Heading1Block{
					RichText: []notion.RichText{{Text: &notion.Text{Content: "우대 사항"}}},
				},
				notion.ParagraphBlock{RichText: []notion.RichText{{Text: &notion.Text{Content: job.PreferredRequirements}}}},
			},
		}

		for retry := 0; retry < 3; retry++ {
			_, err := client.CreatePage(context.Background(), param)
			if err == nil {
				break
			}
			log.Printf("%+v", err)
			log.Printf("failed to add %v page. retry %v", job.Title, retry+1)
			time.Sleep(time.Second * 1)
		}
	}
}
