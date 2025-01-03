package mongo

import (
	"context"
	"fmt"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/pkg/mongo"
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
		fmt.Println(err)
		return models.Movie{}, repository.MapError(err)
	}

	return m, err
}
