package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"crud_ql/auth"
	"crud_ql/graph/model"
	"crud_ql/repository"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	err := auth.SignUp(ctx, *auth_instance, input)
	if err != nil {
		return nil, err
	}

	return db.CreateUser(input), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	return db.UpdateUser(id, input), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, input string) (*model.DeleteUserResponse, error) {
	return db.DeleteUser(input), nil
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	return db.CreateTask(input), nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTaskInput) (*model.Task, error) {
	return db.UpdateTask(id, input), nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*model.DeleteTaskResponse, error) {
	return db.DeleteTask(id), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.GetUsers(), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.GetUser(id), nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return db.GetTasks(), nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	return db.GetTask(id), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) SignUp(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: SignUp - signUp"))
}

var db = repository.Connect()
var auth_instance = auth.Init()
