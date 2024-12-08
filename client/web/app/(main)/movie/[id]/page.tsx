'use client';

import { useParams } from 'next/navigation';
import Image from 'next/image';
import Rating from '../(components)/Rating';
import { TMDB_API } from '@/utils/constants';
import { useEffect, useState } from 'react';
import { formatDate } from '@/lib/utils';
import { fetchMovieCredits, fetchMovieDetail } from '@/utils/apis/movie';
import Loading from '../../loading';
// import CastList from '../(components)/CastList';

interface Genre {
  id: number;
  name: string;
}

interface Movie {
  backdrop_path: string;
  budget: number;
  genres: Genre[];
  id: string;
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

const MovieDetail = () => {
  const params = useParams();
  const { id } = params;
  const [imageSrc, setImageSrc] = useState('/poster-default.svg');
  const [transitioning, setTransitioning] = useState(false);
  const [idMovie, setIdMovie] = useState<number>(0);
  const [movie, setMovie] = useState<Movie | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function pickMovieFields(movie: any) {
    return {
      backdrop_path: movie.backdrop_path,
      budget: movie.budget,
      genres: movie.genres,
      id: movie.id,
      origin_country: movie.origin_country,
      original_language: movie.original_language,
      original_title: movie.original_title,
      overview: movie.overview,
      popularity: movie.popularity,
      poster_path: movie.poster_path,
      release_date: movie.release_date,
      revenue: movie.revenue,
      runtime: movie.runtime,
      status: movie.status,
      tagline: movie.tagline,
      title: movie.title,
      vote_average: movie.vote_average,
      vote_count: movie.vote_count,
    };
  }

  useEffect(() => {
    if (id) {
      const numericId = +id;
      setIdMovie(numericId);
    }
  }, [id]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const movieResponse = await fetchMovieDetail(idMovie);
        const data = pickMovieFields(movieResponse.data);
        setImageSrc(TMDB_API.POSTER(data.poster_path));
        setMovie(data);

        const creditsResponse = await fetchMovieCredits(idMovie);
        console.log(creditsResponse);
      } catch (err) {
        setError('Failed to fetch movie detail');
        console.log(err);
      } finally {
        setLoading(false);
      }
    };

    if (idMovie > 0) {
      fetchData();
    }
  }, [idMovie]);

  useEffect(() => {
    if (loading === true) {
      setTransitioning(true);
      setTimeout(() => {
        setTransitioning(false);
      }, 300);
    }
  }, [loading]);

  function convertMinutes(minutes: number): string {
    const hours = Math.floor(minutes / 60);
    const minutesRemaining = minutes % 60;
    return `${hours}h ${minutesRemaining}m`;
  }

  if (error) return <div>{error}</div>;

  return (
    <div className={`relative transition-opacity duration-300 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
      {movie !== null && !loading && imageSrc != '/poster-default.svg' ? (
        <div className="font-geist-mono">
          <div
            className="relative py-10 px-5 shadow-lg"
            style={{
              backgroundImage: `url(${TMDB_API.POSTER(movie.backdrop_path)})`,
              backgroundSize: '100%',
              backgroundRepeat: 'no-repeat',
              backgroundPosition: 'center top',
            }}
          >
            <div className="absolute z-0 inset-0 bg-black/70"></div>
            <div className="relative z-10 container mx-auto text-white">
              <div className="flex flex-row gap-4">
                <div className="relative w-96 h-[400px] aspect-[2/3] rounded-lg overflow-hidden">
                  <Image
                    src={imageSrc}
                    alt={movie.title}
                    layout="fill"
                    objectFit="contain"
                    onError={() => setImageSrc('poster-default.svg')}
                  />
                </div>

                <div className="flex flex-col justify-around items-start ml-5">
                  <div>
                    <h1 className="text-3xl font-bold">{movie.original_title}</h1>
                    <div className="flex flex-row gap-6 text-sm">
                      <div>{formatDate(movie.release_date)}</div>
                      <div>●</div>
                      <div>{movie.genres.map((genre) => genre.name).join(', ')}</div>
                      <div>●</div>
                      <div>{convertMinutes(movie.runtime)}</div>
                    </div>
                  </div>
                  <div className="flex flex-row items-center gap-2">
                    <div className="relative flex gap-2 items-center">
                      <Rating rating={movie.vote_average} />
                    </div>
                    <div className="flex flex-col text-sm font-bold">
                      <div>User</div>
                      <div>score</div>
                    </div>
                  </div>
                  <p className="text-sm italic">{movie.tagline}</p>
                  <div>
                    <h1 className="text-md font-bold mb-1">Overview</h1>
                    <p className="text-sm max-w-3xl">{movie.overview}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="mt-6">
            <h2 className="text-2xl font-semibold mb-4">Cast</h2>
            {/* <CastList cast={movie.cast} /> */}
          </div>
        </div>
      ) : (
        <div className="font-geist-mono">
          <div
            className="relative py-10 px-5 shadow-lg"
            style={{
              backgroundImage: `url(${TMDB_API.POSTER(imageSrc)})`,
              backgroundSize: '100%',
              backgroundRepeat: 'no-repeat',
              backgroundPosition: 'center top',
            }}
          >
            <div className="absolute z-0 inset-0 bg-black/70"></div>
            <div className="relative z-10 container mx-auto text-white">
              <div className="flex flex-row gap-4">
                <div className="relative w-96 h-[400px] aspect-[2/3] rounded-lg overflow-hidden">
                  <Image
                    src={imageSrc}
                    alt={'loading...'}
                    layout="fill"
                    objectFit="contain"
                    onError={() => setImageSrc('poster-default.svg')}
                  />
                </div>
                <div className="flex flex-col justify-center items-center ml-5">
                  <Loading />
                </div>
              </div>
            </div>
          </div>

          <div className="mt-6">
            <h2 className="text-2xl font-semibold mb-4">Cast</h2>
            {/* <CastList cast={movie.cast} /> */}
          </div>
        </div>
      )}
    </div>
  );
};

export default MovieDetail;
