package tvshow

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOneTVShow(ctx context.Context, tvID int64) (models.TVShow, error)
	ListTVShows(ctx context.Context, input ListTVsOptions) ([]models.TVShow, error)
	CountTVShows(ctx context.Context, input ListTVsOptions) (int64, error)
	GetUpcomingTVShows(ctx context.Context, input GetUpcomingTVsOptions) ([]models.TVSummary, error)
	CountUpcomingTVShows(ctx context.Context, input GetUpcomingTVsOptions) (int64, error)
	GetOnTheAirTVShows(ctx context.Context, input GetOnTheAirTVsOptions) ([]models.TVSummary, error)
	CountOnTheAirTVShows(ctx context.Context, input GetOnTheAirTVsOptions) (int64, error)
	GetAiringTodayTVShows(ctx context.Context, input GetAiringTodayTVsOptions) ([]models.TVSummary, error)
	CountAiringTodayTVShows(ctx context.Context, input GetAiringTodayTVsOptions) (int64, error)
	GetTopRatedTVShows(ctx context.Context, input GetTopRatedTVsOptions) ([]models.TVSummary, error)
	CountTopRatedTVShows(ctx context.Context, input GetTopRatedTVsOptions) (int64, error)
	GetPopularTVShows(ctx context.Context, input GetPopularTVsOptions) ([]models.TVSummary, error)
	CountPopularTVShows(ctx context.Context, input GetPopularTVsOptions) (int64, error)
	GetTVShowGenres(ctx context.Context) ([]models.TVGenre, error)
}
