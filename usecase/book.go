package usecase

import (
	"context"
	"fmt"
	"myapp/graph/model"
)

func BookCreate(ctx context.Context, input model.NewBook) (*model.Book, error) {
	book := &model.Book{
		ID:       DummyBooks[len(DummyBooks)-1].ID + 1,
		Title:    &input.Title,
		AuthorID: input.AuthorID,
	}

	DummyBooks = append(DummyBooks, book)

	return book, nil
}

func BookGetAll(ctx context.Context) ([]*model.Book, error) {
	fmt.Println("SELECT * FROM books")

	return DummyBooks, nil
}

func BookGetByID(ctx context.Context, id int) (*model.Book, error) {
	fmt.Println("SELECT * FROM books WHERE id = ", id)

	for _, book := range DummyBooks {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, nil
}

func BookGetByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error) {
	fmt.Println("SELECT * FROM books WHERE author_id = ", authorID)

	books := []*model.Book{}
	for _, book := range DummyBooks {
		if book.AuthorID == authorID {
			books = append(books, book)
		}
	}

	return books, nil
}
