import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '../constants';

export const fetchMovieDetail = (id: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_DETAILS(id));
};

export const fetchMovieCredits = (id: number) => {
  return axiosInstance.get(TMDB_API.MOVIE_CREDITS(id));
};
