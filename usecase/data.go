package usecase

import (
	"myapp/graph/model"
)

var (
	firstName1   = "Osamu"
	lastName1    = "Dazai"
	title1       = "Overlord IV"
	DummyAuthors = []*model.Author{
		{
			ID:        1,
			FirstName: &firstName1,
			LastName:  &lastName1,
		},
	}
	DummyBooks = []*model.Book{
		{
			ID:       1,
			Title:    &title1,
			AuthorID: 1,
		},
	}
)
