package grapqhlconfig

import (
	gql "graphql_api/graphQL"
	graphqldirectives "graphql_api/graphQL/directives"
	graphqlresolvers "graphql_api/graphQL/resolvers"
)

/*
	Creates and returns the GraphQL configuration
*/
func GetConfig() gql.Config {
	config := gql.Config{Resolvers: &graphqlresolvers.Resolver{}}

	config.Directives.IsSignedIn = graphqldirectives.IsSignedIn

	return config
}
