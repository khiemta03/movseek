'use client';

import { useEffect, useState } from 'react';
import { Carousel, CarouselContent, CarouselItem, CarouselPrevious, CarouselNext } from '@/components/ui/carousel';
import Autoplay from 'embla-carousel-autoplay';
import { fetchTrending } from '@/apis/trending';
import { Video } from '@/models/movie-detail-types';
import { fetchMovieVideos } from '@/apis/movie';
import { selectPreferredVideo } from '@/utils/util-functions/detail-page';
import TrailerCard from '@/components/main/trailer-card';
import TrailerCardDummy from '@/components/main/trailer-card-dummy';

interface LatestTrailerListProps {
  changeThumbnail: (thumbnail: string) => void;
  onPlayVideo: (url: string) => void;
}

const LatestTrailerList: React.FC<LatestTrailerListProps> = ({ changeThumbnail, onPlayVideo }) => {
  const [transitioning, setTransitioning] = useState(false);
  const [trailers, setTrailers] = useState<{ trailer: Video; thumbnail: string; title: string; id: number }[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchTrendingMovies = async () => {
      try {
        setLoading(true);
        const moviesResponse = await fetchTrending('movie', 'day');
        const data = (
          await Promise.all(
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            moviesResponse.data.results.map(async (movie: Record<string, any>) => {
              const videoResponse = await fetchMovieVideos(movie.id);
              const trailer = selectPreferredVideo(videoResponse.data.results);
              return {
                trailer,
                thumbnail: movie.poster_path,
                title: movie.title,
                id: movie.id,
              };
            }),
          )
        ).filter((item) => item.trailer);
        setTrailers(data);
      } catch (err) {
        console.log(err);
      } finally {
        setTransitioning(true);
        setTimeout(() => {
          setLoading(false);
          setTransitioning(false);
        }, 500);
      }
    };

    fetchTrendingMovies();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const moviesDummy = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];

  return (
    <div className="relative w-full py-4">
      <div className={`relative transition-opacity duration-500 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
        <Carousel
          plugins={[
            Autoplay({
              delay: 5000,
            }),
          ]}
        >
          <div className="flex space-x-8">
            <CarouselContent>
              {loading
                ? moviesDummy.map((movie, index) => (
                    <CarouselItem
                      key={index}
                      className="basis-80 h-60"
                    >
                      <TrailerCardDummy />
                    </CarouselItem>
                  ))
                : trailers.map((element, index) => (
                    <CarouselItem
                      key={index}
                      className="basis-80 h-60"
                    >
                      <TrailerCard
                        videoUrl={element.trailer.key}
                        thumbnail={element.thumbnail}
                        id={element.id}
                        title={element.title}
                        description={element.trailer.name}
                        changeThumbnail={changeThumbnail}
                        onPlayVideo={onPlayVideo}
                      />
                      {/* <TrailerCardDummy /> */}
                    </CarouselItem>
                  ))}
            </CarouselContent>
            {!loading && (
              <>
                <div className="absolute top-1/2 left-8 transform -translate-y-1/2">
                  <CarouselPrevious className="p-2 bg-gray-800 text-white rounded-full" />
                </div>
                <div className="absolute top-1/2 right-16 transform -translate-y-1/2">
                  <CarouselNext className="p-2 bg-gray-800 text-white rounded-full" />
                </div>
              </>
            )}
          </div>
        </Carousel>
      </div>
    </div>
  );
};

export default LatestTrailerList;
