package person

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Usecase interface {
	GetOnePerson(ctx context.Context, personID int64) (models.Person, error)
	ListPeople(ctx context.Context, input ListPeopleInput) (ListPeopleOutput, error)
	GetTrendingPeople(ctx context.Context, input GetTrendingPeopleInput) (GetTrendingPeopleOutput, error)
	GetPopularPeople(ctx context.Context, input GetPopularPeopleInput) (GetPopularPeopleOutput, error)
}

