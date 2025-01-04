package mongo

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	movieCollection = "movies"
)

func (repo implRepository) getMovieCollection() mongo.Collection {
	return repo.db.Collection(movieCollection)
}

func (repo implRepository) GetOneMovie(ctx context.Context, movieID string) (models.Movie, error) {
	col := repo.getMovieCollection()

	queryFilter, err := repo.buildMovieQuery(movieID)
	if err != nil {
		return models.Movie{}, err
	}

	var m = models.Movie{}
	err = col.FindOne(ctx, queryFilter).Decode(&m)
	if err != nil {
		return models.Movie{}, repository.MapError(err)
	}

	return m, err
}

func (repo implRepository) ListMovies(ctx context.Context, input movie.ListMoviesOptions) ([]models.Movie, error) {
	col := repo.getMovieCollection()

	queryFilter := repo.buildListMoviesQuery(input)

	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))

	cursor, err := col.Find(ctx, queryFilter, findOptions)
	if err != nil {
		return []models.Movie{}, err
	}

	var movies []models.Movie
	err = cursor.All(ctx, &movies)
	if err != nil {
		return []models.Movie{}, err
	}

	return movies, nil
}
