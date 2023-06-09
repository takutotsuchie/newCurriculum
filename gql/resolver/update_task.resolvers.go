package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"newCurriculum/db"
	"newCurriculum/gql/model"
)

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.NewTask) (string, error) {
	return db.UpdateTask(ctx, input)
}
