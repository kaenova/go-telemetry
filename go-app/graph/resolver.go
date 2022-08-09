package graph

import (
	"github.com/kaenova/go-elastic-kibana/go-app/graph/model"
	"github.com/kaenova/go-elastic-kibana/go-app/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Logger *logger.Logger

	userIdIncrement  int
	todosIdIncrement int

	users []*model.User
	todos []*model.Todo
}

func (r *Resolver) searchUserById(id int) model.User {
	var user model.User
	for _, v := range r.users {
		if v.ID == id {
			user = *v
		}
	}
	return user
}
