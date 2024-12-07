'use client';

import MovieCard from './MovieCard';
import { Carousel, CarouselContent, CarouselItem, CarouselPrevious, CarouselNext } from './ui/carousel';
import Autoplay from 'embla-carousel-autoplay';

interface MovieListProps {
  category: string;
}

const MovieList = ({ category }: MovieListProps) => {
  // Giả sử bạn có dữ liệu mock hoặc gọi API
  const movies = [
    {
      id: 1,
      title: 'A miraculous conception. A merciless king.',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
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
      rating: 7.0,
    },
    {
      id: 4,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 5,
      title: 'Movie 2',
      poster: 'https://th.bing.com/th/id/OIP.QblrzVh2BqODKiAlvUCMqAHaKf?w=188&h=266&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 6,
      title: 'Movie 3',
      poster: 'https://th.bing.com/th/id/OIP.gwUCRkCrlItYPT7oEALsKAHaLH?w=188&h=282&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
    },
    {
      id: 7,
      title: 'Movie 1',
      poster: 'https://th.bing.com/th/id/OIP.Gx-zcrqyJsQYrr1g14qCEAHaK9?w=188&h=279&c=7&r=0&o=5&dpr=2&pid=1.7',
      releaseDate: '2024-12-05',
      rating: 7.0,
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
              <CarouselItem key={movie.id} className="basis-40 h-80">
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
    </div>
  );
};

export default MovieList;
