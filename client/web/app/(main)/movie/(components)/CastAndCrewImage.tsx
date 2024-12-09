import React, { useState } from 'react';
import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';

interface CastAndCrewImageProps {
  image: string;
  name: string;
}

const CastAndCrewImage: React.FC<CastAndCrewImageProps> = ({ image, name }) => {
  const [imageSrc, setImageSrc] = useState(TMDB_API.POSTER(image));

  return (
    <div className="relative w-20 h-20 rounded-full overflow-hidden shadow-md">
      <Image
        className=" hover:cursor-pointer"
        src={imageSrc}
        alt={name}
        layout="fill"
        objectFit="cover"
        onError={() => setImageSrc('/poster-default.svg')}
      />
    </div>
  );
};

export default CastAndCrewImage;
