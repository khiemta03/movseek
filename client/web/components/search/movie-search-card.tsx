import { formatDate } from '@/lib/utils';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import Image from 'next/image';
import Link from 'next/link';
import { TMDB_API } from '@/utils/constants';
import { Movie } from '@/models/search-types';
import { useEffect, useState } from 'react';

interface MovieSearchCardProps {
  movie: Movie;
}

const MovieSearchCard: React.FC<MovieSearchCardProps> = ({ movie }) => {
  const [imageSrc, setImageSrc] = useState(TMDB_API.POSTER(movie.poster_path));

  useEffect(() => {
    setImageSrc(TMDB_API.POSTER(movie.poster_path));
  }, [movie]);

  return (
    <div className="relative w-full h-80 font-geist overflow-hidden bg-gray-100 rounded-lg shadow hover:shadow-lg transition">
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger>
            <Link href={`/movie/${movie.id}`}>
              <Image
                src={movie.poster_path ? imageSrc : '/poster-default.svg'}
                alt={movie.title}
                className="h-60 object-cover hover:cursor-pointer"
                objectFit="cover"
                width={400}
                height={400 * 1.618}
                onError={() => setImageSrc('/poster-default.svg')}
                priority
              />
            </Link>
          </TooltipTrigger>
          <TooltipContent>
            <p>{movie.title}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
      <div className="absolute text-center bottom-0 left-0 right-0 bg-gradient-to-t from-stone-900 to-transparent px-2 py-5">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger>
              <Link href={`/movie/${movie.id}`}>
                <h3 className="text-sm text-white font-bold line-clamp-1 hover:text-primary hover:cursor-pointer">
                  {movie.title}
                </h3>
              </Link>
            </TooltipTrigger>
            <TooltipContent>
              <p>{movie.title}</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
        <p className="text-xs italic mt-1 text-white">
          {movie.release_date != '' && movie.release_date != null ? formatDate(movie.release_date) : ''}
        </p>
      </div>
    </div>
  );
};

export default MovieSearchCard;
