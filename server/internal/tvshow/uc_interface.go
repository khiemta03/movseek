package tvshow

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Usecase interface {
	GetOneTVShow(ctx context.Context, tvID int64) (models.TVShow, error)
	ListTVShows(ctx context.Context, input ListTVsInput) (ListTVsOutput, error)
	GetUpcomingTVShows(ctx context.Context, input GetUpcomingTVsInput) (GetUpcomingTVsOutput, error)
	GetTrendingTVShows(ctx context.Context, input GetTrendingTVsInput) (GetTrendingTVsOutput, error)
	GetTopRatedTVShows(ctx context.Context, input GetTopRatedTVsInput) (GetTopRatedTVsOutput, error)
	GetPopularTVShows(ctx context.Context, input GetPopularTVsInput) (GetPopularTVsOutput, error)
	GetTVShowGenres(ctx context.Context) (GetTVGenresOutput, error)
}
