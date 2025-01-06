import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchTVPopular = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_POPULAR(page));
};

export const fetchTVTopRated = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_TOP_RATED(page));
};

export const fetchTVNowPlaying = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_NOW_PLAYING(page));
};

export const fetchTVUpcoming = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_UPCOMING(page));
};

export const fetchGenresTV = () => {
  return axiosInstance.get(TMDB_API.GENRES_TV());
};
