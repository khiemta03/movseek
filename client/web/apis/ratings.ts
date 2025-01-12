import axiosInstance from '@/lib/axios.config';
import { TMDB_API } from '@/utils/constants';

export const addRating = (
  avatar: string,
  media_id: number,
  rating: number,
  type: 'movie' | 'tv_show',
  user_id: string,
  username: string,
) => {
  return axiosInstance.post(TMDB_API.ADD_RATING(), {
    avatar,
    rating,
    media_id,
    type,
    user_id,
    username,
  });
};

export const getRatingsByMedia = (media_id: number) => {
  return axiosInstance.get(TMDB_API.GET_RATINGS_BY_MEDIA(media_id));
};

export const getRatingsByUser = (user_id: string) => {
  return axiosInstance.get(TMDB_API.GET_RATINGS_BY_USER(user_id));
};

export const updateRating = (
  avatar: string,
  rating: number,
  media_id: number,
  type: 'movie' | 'tv_show',
  user_id: string,
  username: string,
) => {
  return axiosInstance.put(TMDB_API.UPDATE_RATINGS(user_id), {
    avatar,
    rating,
    media_id,
    type,
    userID: user_id,
    username,
  });
};

export const deleteComment = (user_id: string, media_id: number, type: string) => {
  return axiosInstance.delete(TMDB_API.DELETE_RATINGS(user_id, media_id, type));
};
