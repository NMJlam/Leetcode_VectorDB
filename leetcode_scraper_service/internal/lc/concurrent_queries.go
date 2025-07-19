package lc

import (
	"context"
	"fmt"
	"log"

	"github.com/lamned/leetcode-scraper-service/internal/client"
)

type QuestionData struct {
	Slug string 
	Description string 
}



func Get_Slugs() chan QuestionData {
	slugs := make(chan QuestionData)

	client := client.GraphQL[SlugResponse]{}
	client.Initialise("https://leetcode.com/graphql")

	limit := 50 
	skip := 0  

	variables := map[string]any {
		"limit": &limit, 
		"skip": &skip, 
	}

	go func(){
		defer close(slugs)

		for {
			response, err := client.Query(context.Background(), 
										SlugQuery, 
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

func Get_Problem_Description(slugs chan QuestionData) chan QuestionData {
	descriptions := make(chan QuestionData)	

	client := 

	go func(){
		defer close(descriptions)

		for slug := range slugs {

		}
	}
}
