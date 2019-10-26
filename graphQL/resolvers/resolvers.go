package graphqlresolvers

import (
	"context"
	gql "graphql_api/graphQL"
	"graphql_api/models"
	authservices "graphql_api/services/auth"
)

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignUp(ctx context.Context, input models.User) (*models.JWT, error) {
	jwt, err := authservices.SignUp(input)

	if err != nil {
		return &models.JWT{}, err
	}

	return &models.JWT{Token: jwt}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.Credentials) (*models.JWT, error) {
	jwt, err := authservices.Login(input)

	if err != nil {
		return &models.JWT{}, err
	}

	return &models.JWT{Token: jwt}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Protected(ctx context.Context) (*models.Message, error) {
	return &models.Message{"Protected query"}, nil
}
