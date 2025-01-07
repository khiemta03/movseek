package movie

import (
	"context"

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

	return tvshow.ListTVsOutput{TVShows: tvshows}, nil
}

func (uc implUsecase) GetUpcomingTVShows(ctx context.Context, input tvshow.GetUpcomingTVsInput) (tvshow.GetUpcomingTVsOutput, error) {
	tvshows, err := uc.repo.GetUpcomingTVShows(ctx, tvshow.GetUpcomingTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetUpcomingTVsOutput{}, err
	}

	return tvshow.GetUpcomingTVsOutput{TVShows: tvshows}, nil
}

func (uc implUsecase) GetTrendingTVShows(ctx context.Context, input tvshow.GetTrendingTVsInput) (tvshow.GetTrendingTVsOutput, error) {
	tvshows, err := uc.repo.GetTrendingTVShows(ctx, tvshow.GetTrendingTVsOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return tvshow.GetTrendingTVsOutput{}, err
	}

	return tvshow.GetTrendingTVsOutput{TVShows: tvshows}, nil
}

func (uc implUsecase) GetTopRatedTVShows(ctx context.Context, input tvshow.GetTopRatedTVsInput) (tvshow.GetTopRatedTVsOutput, error) {
	tvshows, err := uc.repo.GetTopRatedTVShows(ctx, tvshow.GetTopRatedTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetTopRatedTVsOutput{}, err
	}

	return tvshow.GetTopRatedTVsOutput{TVShows: tvshows}, nil
}

func (uc implUsecase) GetPopularTVShows(ctx context.Context, input tvshow.GetPopularTVsInput) (tvshow.GetPopularTVsOutput, error) {
	tvshows, err := uc.repo.GetPopularTVShows(ctx, tvshow.GetPopularTVsOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return tvshow.GetPopularTVsOutput{}, err
	}

	return tvshow.GetPopularTVsOutput{TVShows: tvshows}, nil
}

func (uc implUsecase) GetTVShowGenres(ctx context.Context) (tvshow.GetTVGenresOutput, error) {
	genres, err := uc.repo.GetTVShowGenres(ctx)
	if err != nil {
		return tvshow.GetTVGenresOutput{}, err
	}

	return tvshow.GetTVGenresOutput{Genres: genres}, nil
}
