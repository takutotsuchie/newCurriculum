package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"newCurriculum/db"
	"newCurriculum/gql"
	"newCurriculum/gql/model"
)

// OnLimit is the resolver for the onLimit field.
func (r *subscriptionResolver) OnLimit(ctx context.Context, input model.Limit) (<-chan string, error) {
	return db.OnLimit(ctx, input)
}

// Subscription returns gql.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gql.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
