import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchSearchMovie = (query: string, page: number) => {
  return axiosInstance.get(TMDB_API.SEARCH_MOVIE(query, page));
};

export const fetchSearchTV = (query: string, page: number) => {
  return axiosInstance.get(TMDB_API.SEARCH_TV(query, page));
};

export const fetchSearchPerson = (query: string, page: number) => {
  return axiosInstance.get(TMDB_API.SEARCH_PERSON(query, page));
};
