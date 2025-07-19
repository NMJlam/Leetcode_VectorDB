package client  

import (
	"fmt"
	"context"

	"github.com/machinebox/graphql"
)

type GraphQL struct {
	URL string  
	Client *graphql.Client	
}	

func (g *GraphQL) Initialise(URL string) {
	g.URL = URL
	g.Client = graphql.NewClient(URL) 
}

func (g * GraphQL) Query (context context.Context, queryString string, response any, variables map[string]any) error {
	request := graphql.NewRequest(queryString)	

	for key, value := range variables {
		request.Var(key, value)
	}

	err := g.Client.Run(context, request, response)

	if err != nil {
		return fmt.Errorf("Error making GraphQL request: %v", err)
	}
	return nil 
}

