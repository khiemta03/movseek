package mongo

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/internal/tvshow"
	"github.com/tmplam/movseek/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	tvshowCollection      = "tv_shows"
	upcomingCollection    = "tv_upcoming"
	onTheAirCollection    = "tv_on_the_air"
	airingTodayCollection = "tv_airing_today"
	topRatedCollection    = "tv_top_rated"
	popularCollection     = "tv_popular"
	genresCollection      = "tv_genres"
)

func (repo implRepository) getTVShowCollection() mongo.Collection {
	return repo.db.Collection(tvshowCollection)
}

func (repo implRepository) getUpcomingCollection() mongo.Collection {
	return repo.db.Collection(upcomingCollection)
}

func (repo implRepository) getOnTheAirCollection() mongo.Collection {
	return repo.db.Collection(onTheAirCollection)
}

func (repo implRepository) getAiringTodayCollection() mongo.Collection {
	return repo.db.Collection(airingTodayCollection)
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

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, input.Query)
	if err != nil {
		return []models.TVShow{}, err
	}

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

func (repo implRepository) CountTVShows(ctx context.Context, input tvshow.ListTVsOptions) (int64, error) {
	col := repo.getTVShowCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, input.Query)
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo implRepository) GetUpcomingTVShows(ctx context.Context, input tvshow.GetUpcomingTVsOptions) ([]models.TVSummary, error) {
	col := repo.getUpcomingCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return []models.TVSummary{}, err
	}

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
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

func (repo implRepository) CountUpcomingTVShows(ctx context.Context, input tvshow.GetUpcomingTVsOptions) (int64, error) {
	col := repo.getUpcomingCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo implRepository) GetOnTheAirTVShows(ctx context.Context, input tvshow.GetOnTheAirTVsOptions) ([]models.TVSummary, error) {
	col := repo.getOnTheAirCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return []models.TVSummary{}, err
	}

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
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

func (repo implRepository) GetAiringTodayTVShows(ctx context.Context, input tvshow.GetAiringTodayTVsOptions) ([]models.TVSummary, error) {
	col := repo.getAiringTodayCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return []models.TVSummary{}, err
	}

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
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

func (repo implRepository) CountOnTheAirTVShows(ctx context.Context, input tvshow.GetOnTheAirTVsOptions) (int64, error) {
	col := repo.getOnTheAirCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo implRepository) CountAiringTodayTVShows(ctx context.Context, input tvshow.GetAiringTodayTVsOptions) (int64, error) {
	col := repo.getAiringTodayCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo implRepository) GetTopRatedTVShows(ctx context.Context, input tvshow.GetTopRatedTVsOptions) ([]models.TVSummary, error) {
	col := repo.getTopRatedCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return []models.TVSummary{}, err
	}

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
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

func (repo implRepository) CountTopRatedTVShows(ctx context.Context, input tvshow.GetTopRatedTVsOptions) (int64, error) {
	col := repo.getTopRatedCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (repo implRepository) GetPopularTVShows(ctx context.Context, input tvshow.GetPopularTVsOptions) ([]models.TVSummary, error) {
	col := repo.getPopularCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return []models.TVSummary{}, err
	}

	findOptions := repo.buildGetTVShowFindOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
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

func (repo implRepository) CountPopularTVShows(ctx context.Context, input tvshow.GetPopularTVsOptions) (int64, error) {
	col := repo.getPopularCollection()

	queryFilter, err := repo.buildListTVShowsQuery(input.Filter, "")
	if err != nil {
		return 0, err
	}

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return count, nil
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
