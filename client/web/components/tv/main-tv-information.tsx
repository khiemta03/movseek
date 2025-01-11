import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';
import { formatDate } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Heart, List } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import type { Cast, Credits, Crew } from '@/models/movie-detail-types';
import { getCrewByJob } from '@/utils/util-functions/detail-page';
import { useState } from 'react';
import Rating from '@/components/movie/rating';
import { TV } from '@/models/tv-detail-types';

interface MainTVInformationProps {
  tv: TV;
  creadits: Credits;
  toggleVideo: () => void;
  hasTrailer: boolean;
}

const MainTVInformation: React.FC<MainTVInformationProps> = ({ tv, creadits, toggleVideo, hasTrailer }) => {
  const [imageSrc, setImageSrc] = useState(TMDB_API.POSTER(tv.poster_path));

  return (
    <div
      className="relative py-10 px-5 shadow-lg"
      style={{
        backgroundImage: `url(${TMDB_API.POSTER(tv.backdrop_path)})`,
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
              src={tv.poster_path ? imageSrc : '/poster-default.svg'}
              alt={tv.name}
              layout="fill"
              objectFit="contain"
              onError={() => setImageSrc('/poster-default.svg')}
            />
          </div>

          <div className="flex flex-col justify-around items-start ml-10">
            <div>
              <h1 className="text-3xl font-bold">{tv.original_name}</h1>
              <div className="flex flex-row gap-6 text-sm">
                {tv.first_air_date && <div>{formatDate(tv.first_air_date)}</div>}
                {tv.genres.length > 0 && (
                  <>
                    <div>●</div>
                    <div>{tv.genres.map((genre) => genre.name).join(', ')}</div>
                  </>
                )}
              </div>
            </div>
            <div className="flex flex-row items-center gap-2">
              <div className="relative">
                <Rating rating={tv.vote_average} />
              </div>
              <div className="flex flex-col text-sm font-bold mr-4">
                <div>User</div>
                <div>score</div>
              </div>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button variant="secondary" size="icon" className="rounded-full p-6 mr-4">
                      <Heart className="text-red-500 font-bold" />
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>Mark as favorite</TooltipContent>
                </Tooltip>
              </TooltipProvider>

              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Button variant="secondary" size="icon" className="rounded-full p-6 mr-4">
                      <List />
                    </Button>
                  </TooltipTrigger>
                  <TooltipContent>Add to your watchlist</TooltipContent>
                </Tooltip>
              </TooltipProvider>

              {hasTrailer && (
                <Button onClick={toggleVideo} className="border" variant="ghost">
                  ▶ Play trailer
                </Button>
              )}
            </div>
            <p className="text-sm italic">{tv.tagline}</p>
            {tv.overview && (
              <div>
                <h1 className="text-md font-bold mb-1">Overview</h1>
                <p className="text-sm max-w-3xl">{tv.overview}</p>
              </div>
            )}
            <div className="flex flex-row justify-between w-full">
              {getCrewByJob(creadits.crew, 'Directing', 'Director').length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Driector</h1>
                  {getCrewByJob(creadits.crew, 'Directing', 'Director').map((crew: Crew) => (
                    <div key={crew.id} className="text-sm">
                      ● {crew.name}
                    </div>
                  ))}
                </div>
              )}
              {getCrewByJob(creadits.crew, 'Writing', 'Writer').length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Writer</h1>
                  {getCrewByJob(creadits.crew, 'Writing', 'Writer').map((crew: Crew) => (
                    <div key={crew.id} className="text-sm">
                      ● {crew.name}
                    </div>
                  ))}
                </div>
              )}
              {creadits.cast.slice(0, 3).length > 0 && (
                <div>
                  <h1 className="text-md font-bold mb-1">Top Cast</h1>
                  {creadits.cast.slice(0, 3).map((crew: Cast) => (
                    <div key={crew.id} className="text-sm">
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

export default MainTVInformation;
