import { Credits } from '@/utils/types';
import CastCard from './CastCard';
import { Carousel, CarouselContent, CarouselItem, CarouselPrevious, CarouselNext } from '@/components/ui/carousel';
import Autoplay from 'embla-carousel-autoplay';
import FullCaseAndCrew from './FullCaseAndCrew';
import { useEffect, useState } from 'react';

interface CastListProps {
  credits: Credits;
  isfull: boolean;
}

const CastList: React.FC<CastListProps> = ({ credits, isfull }) => {
  const [transitioning, setTransitioning] = useState(false);

  useEffect(() => {
    setTransitioning(true);
    setTimeout(() => {
      setTransitioning(false);
    }, 500);
  }, [isfull]);

  return (
    <div className={`relative transition-opacity duration-500 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
      {isfull ? (
        <FullCaseAndCrew credits={credits} />
      ) : (
        <Carousel
          plugins={[
            Autoplay({
              delay: 5000,
            }),
          ]}
        >
          <div className="flex space-x-8">
            <CarouselContent>
              {credits.cast.map((actor) => (
                <CarouselItem key={actor.id} className="basis-44 h-80">
                  <CastCard actor={actor} />
                </CarouselItem>
              ))}
            </CarouselContent>
            <div className="absolute top-1/2 left-8 transform -translate-y-1/2">
              <CarouselPrevious className="p-2 bg-gray-800 text-white rounded-full" />
            </div>
            <div className="absolute top-1/2 right-16 transform -translate-y-1/2">
              <CarouselNext className="p-2 bg-gray-800 text-white rounded-full" />
            </div>
          </div>
        </Carousel>
      )}
    </div>
  );
};

export default CastList;
