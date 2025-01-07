package mongo

import (
	"context"
	"fmt"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/internal/tvshow"
	"github.com/tmplam/movseek/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	tvshowCollection       = "tv_shows"
	upcomingCollection     = "tv_upcoming"
	weekTrendingCollection = "tv_trending_week"
	dayTrendingCollection  = "tv_trending_day"
	topRatedCollection     = "tv_top_rated"
	popularCollection      = "tv_popular"
	genresCollection       = "tv_genres"
)

func (repo implRepository) getTVShowCollection() mongo.Collection {
	return repo.db.Collection(tvshowCollection)
}

func (repo implRepository) getUpcomingCollection() mongo.Collection {
	return repo.db.Collection(upcomingCollection)
}

func (repo implRepository) getTrendingCollection(t string) mongo.Collection {
	return repo.db.Collection(fmt.Sprintf("tv_trending_%s", t))
}

func (repo implRepository) getTopRatedCollection() mongo.Collection {
	return repo.db.Collection(topRatedCollection)
}

func (repo implRepository) getPopularCollection() mongo.Collection {
	return repo.db.Collection(popularCollection)
}

func (repo implRepository) getGenresCollection() mongo.Collection {
	return repo.db.Collection(genresCollection)
}

func (repo implRepository) GetOneTVShow(ctx context.Context, tvID int64) (models.TVShow, error) {
	col := repo.getTVShowCollection()

	queryFilter, err := repo.buildTVShowQuery(tvID)
	if err != nil {
		return models.TVShow{}, err
	}

	var t = models.TVShow{}
	err = col.FindOne(ctx, queryFilter).Decode(&t)
	if err != nil {
		return models.TVShow{}, repository.MapError(err)
	}

	return t, err
}

func (repo implRepository) ListTVShows(ctx context.Context, input tvshow.ListTVsOptions) ([]models.TVShow, error) {
	col := repo.getTVShowCollection()

	queryFilter := repo.buildListTVShowsQuery(input)

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
	if err != nil {
		return []models.TVShow{}, err
	}

	var tvshows []models.TVShow
	err = cursor.All(ctx, &tvshows)
	if err != nil {
		return []models.TVShow{}, err
	}

	return tvshows, nil
}

func (repo implRepository) GetUpcomingTVShows(ctx context.Context, input tvshow.GetUpcomingTVsOptions) ([]models.TVSummary, error) {
	col := repo.getUpcomingCollection()

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.TVSummary{}, err
	}

	var tvshows []models.TVSummary
	err = cursor.All(ctx, &tvshows)
	if err != nil {
		return []models.TVSummary{}, err
	}

	return tvshows, nil
}

func (repo implRepository) GetTrendingTVShows(ctx context.Context, input tvshow.GetTrendingTVsOptions) ([]models.TVSummary, error) {
	col := repo.getTrendingCollection(input.Type)

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.TVSummary{}, err
	}

	var tvshows []models.TVSummary
	err = cursor.All(ctx, &tvshows)
	if err != nil {
		return []models.TVSummary{}, err
	}

	return tvshows, nil
}

func (repo implRepository) GetTopRatedTVShows(ctx context.Context, input tvshow.GetTopRatedTVsOptions) ([]models.TVSummary, error) {
	col := repo.getTopRatedCollection()

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.TVSummary{}, err
	}

	var tvshows []models.TVSummary
	err = cursor.All(ctx, &tvshows)
	if err != nil {
		return []models.TVSummary{}, err
	}

	return tvshows, nil
}

func (repo implRepository) GetPopularTVShows(ctx context.Context, input tvshow.GetPopularTVsOptions) ([]models.TVSummary, error) {
	col := repo.getPopularCollection()

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.TVSummary{}, err
	}

	var tvshows []models.TVSummary
	err = cursor.All(ctx, &tvshows)
	if err != nil {
		return []models.TVSummary{}, err
	}

	return tvshows, nil
}

func (repo implRepository) GetTVShowGenres(ctx context.Context) ([]models.TVGenre, error) {
	col := repo.getGenresCollection()

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return []models.TVGenre{}, err
	}

	var genres []models.TVGenre
	err = cursor.All(ctx, &genres)
	if err != nil {
		return []models.TVGenre{}, err
	}

	return genres, nil
}
