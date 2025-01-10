package movie

import (
	"context"
	"math"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/tvshow"
)

func (uc implUsecase) GetOneTVShow(ctx context.Context, tvID int64) (models.TVShow, error) {
	return uc.repo.GetOneTVShow(ctx, tvID)
}

func (uc implUsecase) ListTVShows(ctx context.Context, input tvshow.ListTVsInput) (tvshow.ListTVsOutput, error) {
	tvshows, err := uc.repo.ListTVShows(ctx, tvshow.ListTVsOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.ListTVsOutput{}, err
	}

	count, err := uc.repo.CountTVShows(ctx, tvshow.ListTVsOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.ListTVsOutput{}, err
	}

	return tvshow.ListTVsOutput{TVShows: tvshows, Pagination: tvshow.Pagination{
		Page:       input.Filter.Page,
		PerPage:    input.Filter.PerPage,
		Total:      int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
	}}, nil
}

func (uc implUsecase) GetUpcomingTVShows(ctx context.Context, input tvshow.GetUpcomingTVsInput) (tvshow.GetUpcomingTVsOutput, error) {
	tvshows, err := uc.repo.GetUpcomingTVShows(ctx, tvshow.GetUpcomingTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetUpcomingTVsOutput{}, err
	}

	count, err := uc.repo.CountUpcomingTVShows(ctx, tvshow.GetUpcomingTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetUpcomingTVsOutput{}, err
	}

	return tvshow.GetUpcomingTVsOutput{
		TVShows: tvshows,
		Pagination: tvshow.Pagination{
			Page:       input.Filter.Page,
			PerPage:    input.Filter.PerPage,
			Total:      int(count),
			TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
		}}, nil
}

func (uc implUsecase) GetOnTheAirTVShows(ctx context.Context, input tvshow.GetOnTheAirTVsInput) (tvshow.GetOnTheAirTVsOutput, error) {
	tvshows, err := uc.repo.GetOnTheAirTVShows(ctx, tvshow.GetOnTheAirTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetOnTheAirTVsOutput{}, err
	}

	count, err := uc.repo.CountOnTheAirTVShows(ctx, tvshow.GetOnTheAirTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetOnTheAirTVsOutput{}, err
	}

	return tvshow.GetOnTheAirTVsOutput{TVShows: tvshows, Pagination: tvshow.Pagination{
		Page:       input.Filter.Page,
		PerPage:    input.Filter.PerPage,
		Total:      int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
	}}, nil
}

func (uc implUsecase) GetAiringTodayTVShows(ctx context.Context, input tvshow.GetAiringTodayTVsInput) (tvshow.GetAiringTodayTVsOutput, error) {
	tvshows, err := uc.repo.GetAiringTodayTVShows(ctx, tvshow.GetAiringTodayTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetAiringTodayTVsOutput{}, err
	}

	count, err := uc.repo.CountAiringTodayTVShows(ctx, tvshow.GetAiringTodayTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetAiringTodayTVsOutput{}, err
	}

	return tvshow.GetAiringTodayTVsOutput{TVShows: tvshows, Pagination: tvshow.Pagination{
		Page:       input.Filter.Page,
		PerPage:    input.Filter.PerPage,
		Total:      int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
	}}, nil
}

func (uc implUsecase) GetTopRatedTVShows(ctx context.Context, input tvshow.GetTopRatedTVsInput) (tvshow.GetTopRatedTVsOutput, error) {
	tvshows, err := uc.repo.GetTopRatedTVShows(ctx, tvshow.GetTopRatedTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetTopRatedTVsOutput{}, err
	}

	count, err := uc.repo.CountTopRatedTVShows(ctx, tvshow.GetTopRatedTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetTopRatedTVsOutput{}, err
	}

	return tvshow.GetTopRatedTVsOutput{TVShows: tvshows, Pagination: tvshow.Pagination{
		Page:       input.Filter.Page,
		PerPage:    input.Filter.PerPage,
		Total:      int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
	}}, nil
}

func (uc implUsecase) GetPopularTVShows(ctx context.Context, input tvshow.GetPopularTVsInput) (tvshow.GetPopularTVsOutput, error) {
	tvshows, err := uc.repo.GetPopularTVShows(ctx, tvshow.GetPopularTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetPopularTVsOutput{}, err
	}

	count, err := uc.repo.CountPopularTVShows(ctx, tvshow.GetPopularTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetPopularTVsOutput{}, err
	}

	return tvshow.GetPopularTVsOutput{TVShows: tvshows, Pagination: tvshow.Pagination{
		Page:       input.Filter.Page,
		PerPage:    input.Filter.PerPage,
		Total:      int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(input.Filter.PerPage))),
	}}, nil
}

func (uc implUsecase) GetTVShowGenres(ctx context.Context) (tvshow.GetTVGenresOutput, error) {
	genres, err := uc.repo.GetTVShowGenres(ctx)
	if err != nil {
		return tvshow.GetTVGenresOutput{}, err
	}

	return tvshow.GetTVGenresOutput{Genres: genres}, nil
}
