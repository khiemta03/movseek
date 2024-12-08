import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '../constants';

export const fetchTrending = (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') => {
  return axiosInstance.get(TMDB_API.TRENDING(mediaType, timeWindow));
};
