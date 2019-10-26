package graphqldirectives

import (
	"context"
	"errors"
	"graphql_api/middlewares"

	"github.com/99designs/gqlgen/graphql"
)

/*
	Returns an error if the user is not set in the context
*/
func IsSignedIn(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if ctx.Value(middlewares.CtxUserFromJwt) == nil {
		return nil, errors.New("you are not authorized")
	}

	return next(ctx)
}
