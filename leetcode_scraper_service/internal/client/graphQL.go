package client  

import (
	"fmt"
	"context"

	"github.com/machinebox/graphql"
)

type GraphQL[T any]struct {
	URL string  
	Client *graphql.Client	
}	

func (g *GraphQL[T]) Initialise(URL string) {
	g.URL = URL
	g.Client = graphql.NewClient(URL) 
}

func (g *GraphQL[T]) Query(context context.Context, queryString string, variables map[string]any) (*T, error) {

	request := graphql.NewRequest(queryString)	

	for key, value := range variables {
		request.Var(key, value)
	}

	var response T 
	err := g.Client.Run(context, request, &response)
	if err != nil {
		return nil, fmt.Errorf("Error making GraphQL request: %v", err)
	}
	return &response, nil 
}

