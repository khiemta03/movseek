package person

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOnePerson(ctx context.Context, personID int64) (models.Person, error)
	ListPeople(ctx context.Context, input ListPeopleOptions) ([]models.Person, error)
	CountPeople(ctx context.Context, input ListPeopleOptions) (int, error)
	GetTrendingPeople(ctx context.Context, input GetTrendingPeopleOptions) ([]models.PersonSummary, error)
	CountTrendingPeople(ctx context.Context, input GetTrendingPeopleOptions) (int, error)
	GetPopularPeople(ctx context.Context, input GetPopularPeopleOptions) ([]models.PersonSummary, error)
	CountPopularPeople(ctx context.Context, input GetPopularPeopleOptions) (int, error)
}
