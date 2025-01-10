package movie

import (
	"context"
	"math"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/person"
)

func (uc implUsecase) GetOnePerson(ctx context.Context, personID int64) (models.Person, error) {
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

	total, err := uc.repo.CountPeople(ctx, person.ListPeopleOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return person.ListPeopleOutput{}, err
	}

	return person.ListPeopleOutput{
		People: people,
		Pagination: person.Pagination{
			Page:         input.Filter.Page,
			PerPage:      input.Filter.PerPage,
			TotalResults: total,
			TotalPages:   int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetTrendingPeople(ctx context.Context, input person.GetTrendingPeopleInput) (person.GetTrendingPeopleOutput, error) {
	people, err := uc.repo.GetTrendingPeople(ctx, person.GetTrendingPeopleOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return person.GetTrendingPeopleOutput{}, err
	}

	total, err := uc.repo.CountTrendingPeople(ctx, person.GetTrendingPeopleOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return person.GetTrendingPeopleOutput{}, err
	}

	return person.GetTrendingPeopleOutput{
		People: people,
		Pagination: person.Pagination{
			Page:         input.Filter.Page,
			PerPage:      input.Filter.PerPage,
			TotalResults: total,
			TotalPages:   int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetPopularPeople(ctx context.Context, input person.GetPopularPeopleInput) (person.GetPopularPeopleOutput, error) {
	people, err := uc.repo.GetPopularPeople(ctx, person.GetPopularPeopleOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return person.GetPopularPeopleOutput{}, err
	}

	total, err := uc.repo.CountPopularPeople(ctx, person.GetPopularPeopleOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return person.GetPopularPeopleOutput{}, err
	}

	return person.GetPopularPeopleOutput{
		People: people,
		Pagination: person.Pagination{
			Page:         input.Filter.Page,
			PerPage:      input.Filter.PerPage,
			TotalResults: total,
			TotalPages:   int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetMovieCredits(ctx context.Context, personID int64) (models.MovieCredit, error) {
	movieCredits, err := uc.repo.GetOnePerson(ctx, personID)
	if err != nil {
		return models.MovieCredit{}, err
	}

	return movieCredits.MovieCredits, nil
}

func (uc implUsecase) GetTVCredits(ctx context.Context, personID int64) (models.TVCredit, error) {
	person, err := uc.repo.GetOnePerson(ctx, personID)
	if err != nil {
		return models.TVCredit{}, err
	}

	return person.TVCredits, nil
}
