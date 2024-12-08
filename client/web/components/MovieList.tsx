'use client';

import { useEffect, useRef, useState } from 'react';
import MovieCard from './MovieCard';
import { Carousel, CarouselContent, CarouselItem, CarouselPrevious, CarouselNext } from './ui/carousel';
import Autoplay from 'embla-carousel-autoplay';

interface MovieListProps {
  category: string;
}

const MovieList = ({ category }: MovieListProps) => {
  const [viewMode, setViewMode] = useState<'carousel' | 'grid'>('carousel');
  const [transitioning, setTransitioning] = useState(false);
  const [backgroundStyle, setBackgroundStyle] = useState({});
  const buttonARef = useRef<HTMLButtonElement>(null);
  const buttonBRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const activeRef = viewMode === 'carousel' ? buttonARef.current : buttonBRef.current;
    if (activeRef) {
      const rect = activeRef.getBoundingClientRect();
      setBackgroundStyle({
        width: `${rect.width}px`,
        height: `${rect.height}px`,
        transform: `translateX(${activeRef.offsetLeft}px)`,
      });
    }
  }, [viewMode]);

  const handleModeChange = (mode: 'carousel' | 'grid') => {
    if (mode !== viewMode) {
      setTransitioning(true);
      setTimeout(() => {
        setViewMode(mode);
        setTransitioning(false);
      }, 500);
    }
  };

  const movies = [
    {
      id: 1,
      title: 'A miraculous conception. A merciless king.',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 8.0,
    },
    {
      id: 2,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 3,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 6.0,
    },
    {
      id: 4,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 5.0,
    },
    {
      id: 5,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 4.0,
    },
    {
      id: 6,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 3.0,
    },
    {
      id: 7,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 0.0,
    },
    {
      id: 8,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 9,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 10,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 11,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 12,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 13,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 14,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 15,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 16,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 17,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 18,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 19,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 20,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 21,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
  ];

  return (
    <div className="relative w-full py-4">
      <div className="flex justify-end">
        <div className="relative items-center bg-white border-2 border-blue-500 mb-6 rounded-full w-fit">
          <div
            className="absolute z-0 top-0 left-0 right-0 bg-blue-500 rounded-full border-white border transition-all duration-300 ease-in-out"
            style={backgroundStyle}
          ></div>
          <button
            ref={buttonARef}
            onClick={() => handleModeChange('carousel')}
            className={`relative z-10 px-4 text-sm py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
              viewMode === 'carousel' ? 'text-white' : 'text-black opacity-60'
            }`}
          >
            View in Carousel
          </button>
          <button
            ref={buttonBRef}
            onClick={() => handleModeChange('grid')}
            className={`relative z-10 px-4 text-sm py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
              viewMode === 'grid' ? 'text-white' : 'text-black opacity-60'
            }`}
          >
            View in Grid
          </button>
        </div>
      </div>
      <div className={`relative transition-opacity duration-500 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
        {viewMode === 'carousel' ? (
          <Carousel
            plugins={[
              Autoplay({
                delay: 5000,
              }),
            ]}
          >
            <div className="flex space-x-8">
              <CarouselContent>
                {movies.map((movie) => (
                  <CarouselItem key={movie.id} className="basis-44 h-80">
                    <MovieCard
                      title={movie.title}
                      poster={movie.poster}
                      releaseDate={movie.releaseDate}
                      rating={movie.rating}
                    />
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
        ) : (
          <div className="grid grid-cols-3 sm:grid-cols-4 md:grid-cols-5 lg:grid-cols-6 xl:grid-cols-8 gap-4">
            {movies.map((movie) => (
              <div key={movie.id} className="h-80">
                <MovieCard
                  title={movie.title}
                  poster={movie.poster}
                  releaseDate={movie.releaseDate}
                  rating={movie.rating}
                />
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default MovieList;
