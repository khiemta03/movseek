import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchTVPopular = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_POPULAR(page));
};

export const fetchTVTopRated = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_TOP_RATED(page));
};

export const fetchTVOnTheAir = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_ON_THE_AIR(page));
};

export const fetchTVAiringToday = (page: number) => {
  return axiosInstance.get(TMDB_API.TV_AIRING_TODAY(page));
};

export const fetchGenresTV = () => {
  return axiosInstance.get(TMDB_API.GENRES_TV());
};
