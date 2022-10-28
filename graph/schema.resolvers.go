package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/usecase"
)

// Books is the resolver for the books field.
func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	return usecase.BookGetByAuthorID(ctx, obj.ID)
}

// Author is the resolver for the author field.
func (r *bookResolver) Author(ctx context.Context, obj *model.Book) (*model.Author, error) {
	return usecase.AuthorGetByID(ctx, obj.AuthorID)
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.Author, error) {
	return usecase.AuthorCreate(ctx, input)
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	return usecase.BookCreate(ctx, input)
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id int) (*model.Book, error) {
	return usecase.BookGetByID(ctx, id)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return usecase.BookGetAll(ctx)
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int) (*model.Author, error) {
	return usecase.AuthorGetByID(ctx, id)
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return usecase.AuthorGetAll(ctx)
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
