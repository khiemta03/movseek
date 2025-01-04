import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const fetchPeoplePopular = (page: number) => {
  return axiosInstance.get(TMDB_API.PEOPLE_POPULAR(page));
};
