import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';
import { formatDate } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Heart, Bookmark } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import type { Cast, Credits, Crew, Movie } from '@/models/movie-detail-types';
import { convertMinutes, getCrewByJob } from '@/utils/util-functions/detail-page';
import { useState } from 'react';
import Rating from '@/components/movie/rating';
import { useToast } from '@/hooks/use-toast';

interface MainMovieInformationProps {
  movie: Movie;
  creadits: Credits;
  toggleVideo: () => void;
  hasTrailer: boolean;
}

const MainMovieInformation: React.FC<MainMovieInformationProps> = ({ movie, creadits, toggleVideo, hasTrailer }) => {
  const { toast } = useToast();
  const [imageSrc, setImageSrc] = useState(TMDB_API.POSTER(movie.poster_path));
  const [favorite, setFavorite] = useState(true);
  const [watchlist, setwatchlist] = useState(true);

  const handlClickFavorite = (favorite: boolean) => {
    //call api here
    setFavorite(!favorite);
    toast({
      title: 'Success',
      description: `${movie.title} was ${favorite ? 'removed from' : 'added to'} your favourite list.`,
      duration: 3000,
      className: 'bg-green-600 text-white border border-gray-200',
    });
  };

  const handlClickWatchlist = (watchlist: boolean) => {
    //call api here
    setwatchlist(!watchlist);
    toast({
      title: 'Success',
      description: `${movie.title} was ${watchlist ? 'removed from' : 'added to'} your watchlist.`,
      duration: 3000,
      className: 'bg-green-600 text-white border border-gray-200',
    });
  };

  return (
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
          <div className="relative h-[450px] aspect-[2/3] rounded-lg overflow-hidden">
            <Image
              src={movie.poster_path ? imageSrc : '/poster-default.svg'}
              alt={movie.title}
              layout="fill"
              objectFit="contain"
              onError={() => setImageSrc('/poster-default.svg')}
            />
          </div>

          <div className="flex flex-col justify-around items-start ml-10">
            <div>
              <h1 className="text-3xl font-bold">{movie.original_title}</h1>
              <div className="flex flex-row gap-6 text-sm">
                {movie.release_date && <div>{formatDate(movie.release_date)}</div>}
                {movie.genres.length > 0 && (
                  <>
                    <div>●</div>
                    <div>{movie.genres.map((genre) => genre.name).join(', ')}</div>
                  </>
                )}
                {movie.runtime != null && movie.runtime > 0 && (
                  <>
                    {''}
                    <div>●</div>
                    <div>{convertMinutes(movie.runtime)}</div>
                  </>
                )}
              </div>
            </div>
            <div className="flex flex-row items-center gap-2">
              <div className="relative">
                <Rating rating={movie.vote_average} />
              </div>
              <div className="flex flex-col text-sm font-bold mr-4">
                <div>User</div>
                <div>score</div>
              </div>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button
                      variant="secondary"
                      size="icon"
                      className={`rounded-full border border-gray-200 ${
                        favorite ? 'bg-red-400 hover:bg-red-500' : 'hover:bg-red-400'
                      } p-6 mr-4`}
                      onClick={() => handlClickFavorite(favorite)}
                    >
                      <Heart className="text-red-400 font-bold fill-white" />
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>{`${favorite ? 'Remove from favorite list' : 'Mark as favorite'}`}</TooltipContent>
                </Tooltip>
              </TooltipProvider>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button
                      variant="secondary"
                      size="icon"
                      className={`rounded-full border border-gray-200 ${
                        watchlist ? 'bg-sky-600 hover:bg-sky-700' : 'hover:bg-sky-600'
                      } p-6 mr-4`}
                      onClick={() => handlClickWatchlist(watchlist)}
                    >
                      <Bookmark className="text-sky-600 font-bold fill-white" />
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>{`${watchlist ? 'Remove from watchlist' : 'Add to your watchlist'}`}</TooltipContent>
                </Tooltip>
              </TooltipProvider>

              {hasTrailer && (
                <Button
                  onClick={toggleVideo}
                  className="border"
                  variant="ghost"
                >
                  ▶ Play trailer
                </Button>
              )}
            </div>
            <p className="text-sm italic">{movie.tagline}</p>
            {movie.overview && (
              <div>
                <h1 className="text-md font-bold mb-1">Overview</h1>
                <p className="text-sm max-w-3xl">{movie.overview}</p>
              </div>
            )}
            <div className="flex flex-row justify-between w-full">
              {getCrewByJob(creadits.crew, 'Directing', 'Director').length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Driector</h1>
                  {getCrewByJob(creadits.crew, 'Directing', 'Director').map((crew: Crew) => (
                    <div
                      key={crew.id}
                      className="text-sm"
                    >
                      ● {crew.name}
                    </div>
                  ))}
                </div>
              )}
              {getCrewByJob(creadits.crew, 'Writing', 'Writer').length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Writer</h1>
                  {getCrewByJob(creadits.crew, 'Writing', 'Writer').map((crew: Crew) => (
                    <div
                      key={crew.id}
                      className="text-sm"
                    >
                      ● {crew.name}
                    </div>
                  ))}
                </div>
              )}
              {creadits.cast.slice(0, 3).length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Top Cast</h1>
                  {creadits.cast.slice(0, 3).map((crew: Cast) => (
                    <div
                      key={crew.id}
                      className="text-sm"
                    >
                      ● {crew.name}
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainMovieInformation;
