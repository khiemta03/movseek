import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchTVDetail = (id: number) => {
  return axiosInstance.get(TMDB_API.TV_DETAILS(id));
};

export const fetchTVCredits = (id: number) => {
  return axiosInstance.get(TMDB_API.TV_CREDITS(id));
};

export const fetchTVKeywords = (id: number) => {
  return axiosInstance.get(TMDB_API.TV_KEYWORDS(id));
};

export const fetchTVVideos = (id: number) => {
  return axiosInstance.get(TMDB_API.TV_VIDEOS(id));
};
