package tvshow

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
)

type Repository interface {
	GetOneTVShow(ctx context.Context, tvID int64) (models.TVShow, error)
	ListTVShows(ctx context.Context, input ListTVsOptions) ([]models.TVShow, error)
	GetUpcomingTVShows(ctx context.Context, input GetUpcomingTVsOptions) ([]models.TVSummary, error)
	GetTrendingTVShows(ctx context.Context, input GetTrendingTVsOptions) ([]models.TVSummary, error)
	GetTopRatedTVShows(ctx context.Context, input GetTopRatedTVsOptions) ([]models.TVSummary, error)
	GetPopularTVShows(ctx context.Context, input GetPopularTVsOptions) ([]models.TVSummary, error)
	GetTVShowGenres(ctx context.Context) ([]models.TVGenre, error)
}
