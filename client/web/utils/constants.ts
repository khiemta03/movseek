export const TMDB_API = {
  TRENDING: (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') =>
    `/trending/${mediaType}/${timeWindow}`,
  MOVIE_DETAILS: (movieId: number) => `/movie/${movieId}`,
  MOVIE_CREDITS: (movieId: number) => `/movie/${movieId}/credits`,
  SEARCH_MOVIE: (query: string) => `/search/movie?query=${query}`,
  POSTER: (poster: string) => `https://image.tmdb.org/t/p/original/${poster}`,
};
