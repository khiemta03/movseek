export const TMDB_API = {
  TRENDING: (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') =>
    `/trending/${mediaType}/${timeWindow}`,
  MOVIE_DETAILS: (movieId: number) => `/movie/${movieId}`,
  MOVIE_CREDITS: (movieId: number) => `/movie/${movieId}/credits`,
  MOVIE_KEYWORDS: (movieId: number) => `/movie/${movieId}/keywords`,
  MOVIE_VIDEOS: (movieId: number) => `/movie/${movieId}/videos`,
  SEARCH_MOVIE: (query: string, page: number) => `/search/movie?query=${query}&page=${page}`,
  SEARCH_TV: (query: string, page: number) => `/search/tv?query=${query}&page=${page}`,
  SEARCH_PERSON: (query: string, page: number) => `/search/person?query=${query}&page=${page}`,
  POSTER: (poster: string) => `https://image.tmdb.org/t/p/original/${poster}`,

  PEOPLE_POPULAR: (page: number) => `/person/popular?page=${page}`,
  PEOPLE_DETAIL: (id: number) => `/person/${id}`,
  PEOPLE_MOVIE_CREDITS: (id: number) => `/person/${id}/movie_credits`,
  PEOPLE_TV_CREDITS: (id: number) => `/person/${id}/tv_credits`,
};
