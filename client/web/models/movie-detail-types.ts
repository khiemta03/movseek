interface Genre {
  id: number;
  name: string;
}

interface Movie {
  backdrop_path: string;
  budget: number;
  genres: Genre[];
  id: string;
  homepage: string;
  origin_country: string[];
  original_language: string;
  original_title: string;
  overview: string;
  popularity: number;
  poster_path: string;
  release_date: string;
  revenue: number;
  runtime: number;
  status: string;
  tagline: string;
  title: string;
  vote_average: number;
  vote_count: number;
}

interface Cast {
  id: number;
  name: string;
  gender: number;
  original_name: string;
  popularity: number;
  profile_path: string;
  cast_id: number;
  character: string;
  credit_id: string;
  order: number;
}

interface Crew {
  id: number;
  name: string;
  gender: number;
  original_name: string;
  popularity: number;
  profile_path: string;
  credit_id: string;
  department: string;
  job: string;
}

interface CrewGroupedByDepartment {
  [department: string]: Crew[];
}

interface Credits {
  cast: Cast[];
  crew: CrewGroupedByDepartment;
}

interface Keyword {
  id: number;
  name: string;
}

export type { Genre, Movie, Cast, Crew, CrewGroupedByDepartment, Credits, Keyword };
