package person

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOnePerson(ctx context.Context, personID int64) (models.Person, error)
	ListPeople(ctx context.Context, input ListPeopleOptions) ([]models.Person, error)
	GetTrendingPeople(ctx context.Context, input GetTrendingPeopleOptions) ([]models.PersonSummary, error)
	GetPopularPeople(ctx context.Context, input GetPopularPeopleOptions) ([]models.PersonSummary, error)
}
