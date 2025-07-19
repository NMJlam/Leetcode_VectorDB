package lc

import (
	"context"
	"log"

	"github.com/lamned/leetcode-scraper-service/internal/client"
)

type QuestionData struct {
	Slug string 
	Description string 
}

type Client struct {
	Client client.GraphQL
}

func CreateLeetcodeScraper() *Client {
	graphqlClient := client.GraphQL{}
	graphqlClient.Initialise("https://leetcode.com/graphql")
	return &Client {
		Client: graphqlClient, 
	}
}

func (lc *Client) Get_Slugs() chan QuestionData {
	slugs := make(chan QuestionData)

	limit := 50 
	skip := 0  

	variables := map[string]any {
		"limit": &limit, 
		"skip": &skip, 
	}

	go func(){
		defer close(slugs)

		for {
			response := SlugResponse{}
			err := lc.Client.Query(context.Background(), 
										SlugQuery,
										&response, 
										variables)

			if err != nil {
				log.Fatalf("Cannot query: %v", err)
			}

			questions := response.ProblemsetQuestionList.Questions
			for _, question := range questions {
				slugs <- QuestionData{
					Slug: question.TitleSlug,
				}
			} 
			
			// we have reached the end of the page if the api response is less that the limit we have set 
			if pageCapacity := len(response.ProblemsetQuestionList.Questions); pageCapacity < limit {
				break 
			}

			skip += limit 
		}
	}()

	return slugs 
} 

