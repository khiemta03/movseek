'use client';

import { useParams } from 'next/navigation';
import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';
import { useEffect, useState } from 'react';
import { fetchMovieCredits, fetchMovieDetail } from '@/utils/apis/movie';
import Loading from '../../loading';
import type { Movie, Credits } from '@/utils/types';
import { pickMovieFields, handleMovieCredits } from '@/utils/util-functions/detail-page';
import CastList from '../(components)/CastList';
import MainMovieInformation from '../(components)/MainMovieInformation';

const MovieDetail = () => {
  const params = useParams();
  const { id } = params;
  const [imageSrc, setImageSrc] = useState('/poster-default.svg');
  const [transitioning, setTransitioning] = useState(false);
  const [idMovie, setIdMovie] = useState<number>(0);
  const [movie, setMovie] = useState<Movie | null>(null);
  const [creadits, setCredits] = useState<Credits | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

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
        setCredits(handleMovieCredits(creditsResponse.data));
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

  if (error) return <div>{error}</div>;

  return (
    <div className={`relative transition-opacity duration-300 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
      {movie !== null && creadits != null && !loading && imageSrc != '/poster-default.svg' ? (
        <div className="font-geist-mono">
          <MainMovieInformation movie={movie} creadits={creadits} />

          <div className="mt-6">
            <h2 className="text-2xl font-semibold mb-4">Cast</h2>
            <CastList credits={creadits} />
          </div>
        </div>
      ) : (
        <div className="font-geist-mono">
          <div
            className="relative py-10 px-5 shadow-lg"
            style={{
              backgroundImage: `url(${TMDB_API.POSTER('/poster-default.svg')})`,
              backgroundSize: '100%',
              backgroundRepeat: 'no-repeat',
              backgroundPosition: 'center top',
            }}
          >
            <div className="absolute z-0 inset-0 bg-black/70"></div>
            <div className="relative z-10 container mx-auto text-white">
              <div className="flex flex-row px-20 gap-4">
                <div className="relative w-96 h-[450px] aspect-[2/3] rounded-lg overflow-hidden">
                  <Image
                    src="/poster-default.svg"
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
