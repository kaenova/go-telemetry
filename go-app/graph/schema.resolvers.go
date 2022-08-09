package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"math/rand"

	"github.com/kaenova/go-elastic-kibana/go-app/graph/generated"
	"github.com/kaenova/go-elastic-kibana/go-app/graph/model"
	"github.com/sirupsen/logrus"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := r.searchUserById(input.UserID)
	todo := model.Todo{
		ID:   r.todosIdIncrement,
		Text: input.Text,
		Done: false,
		User: &user,
	}
	r.todos = append(r.todos, &todo)
	r.todosIdIncrement++
	return &todo, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		ID:   r.userIdIncrement,
		Name: input.Name,
	}
	r.userIdIncrement++
	r.users = append(r.users, &user)
	return &user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	val := rand.Float32()

	if val < 0.4 {
		r.Logger.Log.WithFields(logrus.Fields{
			"function": "Todos",
		}).Error("cannot fetch data")
		return nil, errors.New("cannot fetch data")
	}

	return r.todos, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	val := rand.Float32()

	if val < 0.4 {
		r.Logger.Log.WithFields(logrus.Fields{
			"function": "Users",
		}).Error("cannot fetch data")
		return nil, errors.New("cannot fetch data")
	}

	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
