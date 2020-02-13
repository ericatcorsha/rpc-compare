package graphql

import (
	"context"
	"errors"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Echo(ctx context.Context, input EchoRequest) (*EchoReply, error) {
	if input.Message == "user-error" {
		return nil, errors.New("The User Made an Error")
	}
	return &EchoReply{Message: input.Message}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "v1", nil
}
