export const TMDB_API = {
  TRENDING: (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') =>
    `/trending/${mediaType}/${timeWindow}`,
  MOVIE_DETAILS: (movieId: number) => `/movie/${movieId}`,
  SEARCH_MOVIE: (query: string) => `/search/movie?query=${query}`,
};
