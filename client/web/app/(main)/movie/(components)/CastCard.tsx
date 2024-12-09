import React, { useState } from 'react';
import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';
import { Cast } from '@/utils/types';

interface CastCardProps {
  actor: Cast;
}

const CastCard: React.FC<CastCardProps> = ({ actor }) => {
  const [imageSrc, setImageSrc] = useState(TMDB_API.POSTER(actor.profile_path));

  return (
    <div className="flex flex-col items-center">
      <div className="relative w-20 h-20 rounded-full overflow-hidden shadow-md">
        <Image
          src={imageSrc}
          alt={actor.name}
          layout="fill"
          objectFit="cover"
          onError={() => setImageSrc('/poster-default.svg')}
        />
      </div>
      <p className="text-sm font-medium text-center mt-2">{actor.name}</p>
    </div>
  );
};

export default CastCard;
