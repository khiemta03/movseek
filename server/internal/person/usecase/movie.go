package movie

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/person"
)

func (uc implUsecase) GetOnePerson(ctx context.Context, personID string) (models.Person, error) {
	return uc.repo.GetOnePerson(ctx, personID)
}

func (uc implUsecase) ListPeople(ctx context.Context, input person.ListPeopleInput) (person.ListPeopleOutput, error) {
	people, err := uc.repo.ListPeople(ctx, person.ListPeopleOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return person.ListPeopleOutput{}, err
	}

	return person.ListPeopleOutput{People: people}, nil
}

func (uc implUsecase) GetTrendingPeople(ctx context.Context, input person.GetTrendingPeopleInput) (person.GetTrendingPeopleOutput, error) {
	people, err := uc.repo.GetTrendingPeople(ctx, person.GetTrendingPeopleOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return person.GetTrendingPeopleOutput{}, err
	}

	return person.GetTrendingPeopleOutput{People: people}, nil
}

func (uc implUsecase) GetPopularPeople(ctx context.Context, input person.GetPopularPeopleInput) (person.GetPopularPeopleOutput, error) {
	people, err := uc.repo.GetPopularPeople(ctx, person.GetPopularPeopleOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return person.GetPopularPeopleOutput{}, err
	}

	return person.GetPopularPeopleOutput{People: people}, nil
}
