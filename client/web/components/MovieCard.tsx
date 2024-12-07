import { formatDate } from '@/lib/utils';
import Image from 'next/image';
import React from 'react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';

interface MovieCardProps {
  title: string;
  releaseDate: string;
  poster: string;
  rating: number; // Rating to display in the progress bar
  onDetailsClick: () => void;
}

const MovieCard: React.FC<MovieCardProps> = ({ title, releaseDate, poster, rating, onDetailsClick }) => {
  const width = 100;
  return (
    <div className="relative w-full h-full bg-white font-geist shadow-lg rounded-lg overflow-hidden hover:shadow-xl transition-shadow">
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger>
            <Image
              src={poster}
              alt={title}
              className="w-full h-60 object-cover hover:cursor-pointer"
              width={width}
              height={width * 1.618}
              priority
            />
          </TooltipTrigger>
          <TooltipContent>
            <p>{title}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>

      <div className="absolute bottom-16 left-4 flex items-center justify-center w-10 h-10 rounded-full">
        <div
          className="absolute inset-0 rounded-full"
          style={{
            background: `conic-gradient(#4caf50 ${rating * 10}%, #29432c ${rating * 10}%)`,
          }}
        ></div>
        <div className="absolute inset-1 bg-gray-800 rounded-full"></div>
        <span className="absolute text-white text-xs font-bold">{rating * 10}%</span>
      </div>

      <div className="absolute top-3/4 bottom-0 left-0 right-0 bg-gradient-to-t from-stone-900 to-transparent px-1 py-5">
        <h3 className="text-xs text-white line-clamp-2 hover:text-primary hover:cursor-pointer">{title}</h3>
        <p className="text-xs text-white">{formatDate(releaseDate)}</p>
      </div>
    </div>
  );
};

export default MovieCard;
