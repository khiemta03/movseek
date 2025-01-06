import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchMoviePopular = (page: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_POPULAR(page));
};

export const fetchGenresMovie = () => {
  return axiosInstance.get(TMDB_API.GENRES_MOVIE());
};
