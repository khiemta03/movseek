import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchMoviePopular = (page: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_POPULAR(page));
};

export const fetchMovieTopRated = (page: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_TOP_RATED(page));
};

export const fetchMovieNowPlaying = (page: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_NOW_PLAYING(page));
};

export const fetchMovieUpcoming = (page: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_UPCOMING(page));
};

export const fetchGenresMovie = () => {
  return axiosInstance.get(TMDB_API.GENRES_MOVIE());
};

export const fetchLatestTrailersMovie = () => {
  return axiosInstance.get(TMDB_API.MOVIE_LATEST_TRAILERS());
};
