package movie

import (
	"context"
	"math"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie"
)

func (uc implUsecase) GetOneMovie(ctx context.Context, movieID int64) (models.Movie, error) {
	return uc.repo.GetOneMovie(ctx, movieID)
}

func (uc implUsecase) GetMovieCredits(ctx context.Context, movieID int64) (models.MovieCredits, error) {
	m, err := uc.repo.GetOneMovie(ctx, movieID)
	if err != nil {
		return models.MovieCredits{}, err
	}

	return m.Credits, nil
}

func (uc implUsecase) GetMovieVideos(ctx context.Context, movieID int64) ([]models.Trailer, error) {
	m, err := uc.repo.GetOneMovie(ctx, movieID)
	if err != nil {
		return []models.Trailer{}, err
	}

	return m.Trailers, nil
}

func (uc implUsecase) GetMovieKeywords(ctx context.Context, movieID int64) ([]models.KeyWord, error) {
	m, err := uc.repo.GetOneMovie(ctx, movieID)
	if err != nil {
		return []models.KeyWord{}, err
	}

	return m.Keywords, nil
}

func (uc implUsecase) ListMovies(ctx context.Context, input movie.ListMoviesInput) (movie.ListMoviesOutput, error) {
	total, err := uc.repo.CountMovies(ctx, movie.ListMoviesOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return movie.ListMoviesOutput{}, err
	}

	movies, err := uc.repo.ListMovies(ctx, movie.ListMoviesOptions{
		Query:  input.Query,
		Filter: input.Filter,
	})
	if err != nil {
		return movie.ListMoviesOutput{}, err
	}

	return movie.ListMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetUpcomingMovies(ctx context.Context, input movie.GetUpcomingMoviesInput) (movie.GetUpcomingMoviesOutput, error) {
	total, err := uc.repo.CountUpcomingMovies(ctx, movie.GetUpcomingMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetUpcomingMoviesOutput{}, err
	}

	movies, err := uc.repo.GetUpcomingMovies(ctx, movie.GetUpcomingMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetUpcomingMoviesOutput{}, err
	}

	return movie.GetUpcomingMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetTrendingMovies(ctx context.Context, input movie.GetTrendingMoviesInput) (movie.GetTrendingMoviesOutput, error) {
	total, err := uc.repo.CountTrendingMovies(ctx, movie.GetTrendingMoviesOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return movie.GetTrendingMoviesOutput{}, err
	}

	movies, err := uc.repo.GetTrendingMovies(ctx, movie.GetTrendingMoviesOptions{
		Filter: input.Filter,
		Type:   input.Type,
	})
	if err != nil {
		return movie.GetTrendingMoviesOutput{}, err
	}

	return movie.GetTrendingMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetTopRatedMovies(ctx context.Context, input movie.GetTopRatedMoviesInput) (movie.GetTopRatedMoviesOutput, error) {
	total, err := uc.repo.CountTopRatedMovies(ctx, movie.GetTopRatedMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetTopRatedMoviesOutput{}, err
	}

	movies, err := uc.repo.GetTopRatedMovies(ctx, movie.GetTopRatedMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetTopRatedMoviesOutput{}, err
	}

	return movie.GetTopRatedMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetPopularMovies(ctx context.Context, input movie.GetPopularMoviesInput) (movie.GetPopularMoviesOutput, error) {
	total, err := uc.repo.CountPopularMovies(ctx, movie.GetPopularMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetPopularMoviesOutput{}, err
	}

	movies, err := uc.repo.GetPopularMovies(ctx, movie.GetPopularMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetPopularMoviesOutput{}, err
	}

	return movie.GetPopularMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetNowPlayingMovies(ctx context.Context, input movie.GetNowPlayingMoviesInput) (movie.GetNowPlayingMoviesOutput, error) {
	total, err := uc.repo.CountNowPlayingMovies(ctx, movie.GetNowPlayingMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetNowPlayingMoviesOutput{}, err
	}

	movies, err := uc.repo.GetNowPlayingMovies(ctx, movie.GetNowPlayingMoviesOptions{
		Filter: input.Filter,
	})
	if err != nil {
		return movie.GetNowPlayingMoviesOutput{}, err
	}

	return movie.GetNowPlayingMoviesOutput{
		Movies: movies,
		Pagination: movie.Pagination{
			Page:      input.Filter.Page,
			PerPage:   input.Filter.PerPage,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(input.Filter.PerPage))),
		},
	}, nil
}

func (uc implUsecase) GetMovieGenres(ctx context.Context) (movie.GetMovieGenresOutput, error) {
	genres, err := uc.repo.GetMovieGenres(ctx)
	if err != nil {
		return movie.GetMovieGenresOutput{}, err
	}

	return movie.GetMovieGenresOutput{Genres: genres}, nil
}

func (uc implUsecase) GetLastestTrailer(ctx context.Context, input movie.GetLastestTrailerInput) ([]movie.GetLastestTrailerResponse, error) {
	trailers, err := uc.repo.GetLastestTrailer(ctx, input)
	if err != nil {
		return []movie.GetLastestTrailerResponse{}, err
	}

	return trailers, nil
}
