export const TMDB_API = {
  // TRENDING: (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') =>
  //   `/trending/${mediaType}/${timeWindow}`,

  // MOVIE_DETAILS: (movieId: number) => `/movie/${movieId}`,
  // MOVIE_CREDITS: (movieId: number) => `/movie/${movieId}/credits`,
  // MOVIE_KEYWORDS: (movieId: number) => `/movie/${movieId}/keywords`,
  // MOVIE_VIDEOS: (movieId: number) => `/movie/${movieId}/videos`,

  // TV_DETAILS: (tvId: number) => `/tv/${tvId}`,
  // TV_CREDITS: (tvId: number) => `/tv/${tvId}/credits`,
  // TV_KEYWORDS: (tvId: number) => `/tv/${tvId}/keywords`,
  // TV_VIDEOS: (tvId: number) => `/tv/${tvId}/videos`,

  // SEARCH_MOVIE: (query: string, page: number) => `/search/movie?query=${query}&page=${page}`,
  // SEARCH_TV: (query: string, page: number) => `/search/tv?query=${query}&page=${page}`,
  // SEARCH_PERSON: (query: string, page: number) => `/search/person?query=${query}&page=${page}`,

  // POSTER: (poster: string) => `https://image.tmdb.org/t/p/original/${poster}`,
  // THUMBNAIL: (thumbnail: string) => `https://image.tmdb.org/t/p/w710_and_h400_multi_faces/${thumbnail}`,

  // PEOPLE_POPULAR: (page: number) => `/person/popular?page=${page}`,
  // PEOPLE_DETAIL: (id: number) => `/person/${id}`,
  // PEOPLE_MOVIE_CREDITS: (id: number) => `/person/${id}/movie_credits`,
  // PEOPLE_TV_CREDITS: (id: number) => `/person/${id}/tv_credits`,

  // GENRES_MOVIE: () => `/genre/movie/list`,
  // MOVIE_POPULAR: (page: number) => `/movie/popular?page=${page}`,
  // MOVIE_TOP_RATED: (page: number) => `/movie/top_rated?page=${page}`,
  // MOVIE_NOW_PLAYING: (page: number) => `/movie/now_playing?page=${page}`,
  // MOVIE_UPCOMING: (page: number) => `/movie/upcoming?page=${page}`,

  // GENRES_TV: () => `/genre/tv/list`,
  // TV_POPULAR: (page: number) => `/tv/popular?page=${page}`,
  // TV_TOP_RATED: (page: number) => `/tv/top_rated?page=${page}`,
  // TV_ON_THE_AIR: (page: number) => `/tv/on_the_air?page=${page}`,
  // TV_AIRING_TODAY: (page: number) => `/tv/airing_today?page=${page}`,

  TRENDING: (mediaType: 'all' | 'movie' | 'tv' | 'person', timeWindow: 'day' | 'week') =>
    `/${mediaType}/trending/${timeWindow}?page=1&per_page=18`,

  MOVIE_DETAILS: (movieId: number) => `/movie/${movieId}`,
  MOVIE_CREDITS: (movieId: number) => `/movie/${movieId}/credits`,
  MOVIE_KEYWORDS: (movieId: number) => `/movie/${movieId}/keywords`,
  MOVIE_VIDEOS: (movieId: number) => `/movie/${movieId}/videos`,

  TV_DETAILS: (tvId: number) => `/tvshow/${tvId}`,
  TV_CREDITS: (tvId: number) => `/tvshow/${tvId}/credits`,
  TV_KEYWORDS: (tvId: number) => `/tvshow/${tvId}/keywords`,
  TV_VIDEOS: (tvId: number) => `/tvshow/${tvId}/videos`,
  TV_CREDITS_TMDB: (tvId: number) => `/tv/${tvId}/credits`,
  TV_KEYWORDS_TMDB: (tvId: number) => `/tv/${tvId}/keywords`,
  TV_VIDEOS_TMDB: (tvId: number) => `/tv/${tvId}/videos`,

  SEARCH_MOVIE: (query: string, page: number) => `/search/movie?query=${query}&page=${page}&per_page=18`,
  SEARCH_SPECIFIC_MOVIE: (ids: string, page: number) => `/search/movie?page=${page}&per_page=10&${ids}`,
  SEARCH_TV: (query: string, page: number) => `/search/tv?query=${query}&page=${page}&per_page=18`,
  SEARCH_PERSON: (query: string, page: number) => `/search/person?query=${query}&page=${page}&per_page=18`,

  POSTER: (poster: string) => `https://image.tmdb.org/t/p/original/${poster}`,
  THUMBNAIL: (thumbnail: string) => `https://image.tmdb.org/t/p/w710_and_h400_multi_faces/${thumbnail}`,

  PEOPLE_POPULAR: (page: number) => `/person/popular?page=${page}&per_page=18`,
  PEOPLE_DETAIL: (id: number) => `/person/${id}`,
  PEOPLE_MOVIE_CREDITS: (id: number) => `/person/${id}/credits/movie`,
  PEOPLE_TV_CREDITS: (id: number) => `/person/${id}/credits/tv`,

  GENRES_MOVIE: () => `/movie/genres`,
  MOVIE_POPULAR: (page: number) => `/movie/popular?page=${page}&per_page=18`,
  MOVIE_TOP_RATED: (page: number) => `/movie/top-rated?page=${page}&per_page=18`,
  MOVIE_NOW_PLAYING: (page: number) => `/movie/now-playing?page=${page}&per_page=18`,
  MOVIE_UPCOMING: (page: number) => `/movie/upcoming?page=${page}&per_page=18`,

  MOVIE_LATEST_TRAILERS: () => `/movie/trailer/latest?page=2&per_page=15`,

  GENRES_TV: () => `/tvshow/genres`,
  TV_POPULAR: (page: number) => `/tvshow/popular?page=${page}&per_page=18`,
  TV_TOP_RATED: (page: number) => `/tvshow/top-rated?page=${page}&per_page=18`,
  TV_ON_THE_AIR: (page: number) => `/tvshow/on-the-air?page=${page}&per_page=18`,
  TV_AIRING_TODAY: (page: number) => `/tvshow/airing-today?page=${page}&per_page=18`,

  ADD_SAVED_ITEM: () => `/saved-item/add`,
  REMOVE_SAVED_ITEM: () => `/saved-item/remove`,
  GET_FAVORITE_ITEM: (userId: string) => `/saved-item/${userId}?type=favorite`,
  GET_WATCHLIST_ITEM: (userId: string) => `/saved-item/${userId}?type=watchlist`,
};
