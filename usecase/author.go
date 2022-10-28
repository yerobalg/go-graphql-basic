package usecase

import (
	"context"
	"myapp/graph/model"
	"fmt"
)

func AuthorCreate(ctx context.Context, input model.NewAuthor) (*model.Author, error) {	
	author := &model.Author{
		ID:        DummyAuthors[len(DummyAuthors)-1].ID + 1,
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
	}
	
	DummyAuthors = append(DummyAuthors, author)

	return author, nil
}

func AuthorGetAll(ctx context.Context) ([]*model.Author, error) {
	fmt.Println("SELECT * FROM authors")
	return DummyAuthors, nil
}

func AuthorGetByID(ctx context.Context, id int) (*model.Author, error) {
	fmt.Println("SELECT * FROM authors WHERE id = ", id)

	for _, author := range DummyAuthors {
		if author.ID == id {
			return author, nil
		}
	}

	return nil, nil
}