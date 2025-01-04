interface MovieKnownFor {
  backdrop_path: string;
  id: number;
  title: string;
  original_title: string;
  overview: string;
  poster_path: string;
  adult: boolean;
  genre_ids: number[];
  original_language: string;
  popularity: number;
  release_date: string;
  video: boolean;
  vote_average: number;
  vote_count: number;
  media_type: 'movie';
}

interface TVShowKnownFor {
  backdrop_path: string;
  id: number;
  name: string;
  original_name: string;
  overview: string;
  poster_path: string;
  genre_ids: number[];
  original_language: string;
  popularity: number;
  first_air_date: string;
  vote_average: number;
  vote_count: number;
  media_type: 'tv';
  adult: boolean;
  origin_country: string[];
}

type KnownFor = MovieKnownFor | TVShowKnownFor;

interface PeoplePopular {
  adult: boolean;
  gender: number;
  id: number;
  known_for_department: string;
  name: string;
  original_name: string;
  popularity: number;
  profile_path: string;
  known_for: KnownFor[];
}

interface PeoplePopularResults {
  page: number;
  results: PeoplePopular[];
  total_pages: number;
  total_results: number;
}

interface Cast {
  adult: boolean;
  backdrop_path: string | null;
  genre_ids: number[];
  id: number;
  original_language: string;
  original_title: string;
  overview: string;
  popularity: number;
  poster_path: string;
  release_date: string;
  title: string;
  video: boolean;
  vote_average: number;
  vote_count: number;
  character: string;
  credit_id: string;
  order: number;
}

interface Crew {
  adult: boolean;
  backdrop_path: string | null;
  genre_ids: number[];
  id: number;
  original_language: string;
  original_title: string;
  overview: string;
  popularity: number;
  poster_path: string;
  release_date: string;
  title: string;
  video: boolean;
  vote_average: number;
  vote_count: number;
  credit_id: string;
  department: string;
  job: string;
}

interface PeopleDetail {
  adult: boolean;
  also_known_as: string[];
  biography: string;
  birthday: string;
  deathday: string | null;
  gender: number;
  homepage: string;
  id: number;
  imdb_id: string;
  known_for_department: string;
  name: string;
  place_of_birth: string;
  popularity: number;
  profile_path: string;
  movie_credits: {
    cast: Cast[];
    crew: Crew[];
    id: number;
  };
}

export type { PeoplePopularResults, MovieKnownFor, PeoplePopular, PeopleDetail };
