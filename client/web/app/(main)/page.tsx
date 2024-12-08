'use client';

import MovieList from '@/components/MovieList';
import { useEffect, useRef, useState } from 'react';

export default function Home() {
  const [activeButton, setActiveButton] = useState<'A' | 'B'>('A');
  const [backgroundStyle, setBackgroundStyle] = useState({});
  const buttonARef = useRef<HTMLButtonElement>(null);
  const buttonBRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const activeRef = activeButton === 'A' ? buttonARef.current : buttonBRef.current;
    if (activeRef) {
      const rect = activeRef.getBoundingClientRect();
      setBackgroundStyle({
        width: `${rect.width}px`,
        height: `${rect.height}px`,
        transform: `translateX(${activeRef.offsetLeft}px)`,
      });
    }
  }, [activeButton]);

  return (
    <div className="text-center italic font-geist-mono">
      <div className="bg-gradient-to-r from-purple-500 via-cyan-500 to-purple-500 py-10 px-5 shadow-lg text-center">
        <div className="container mx-auto">
          <h1 className="text-5xl font-bold bg-gradient-to-l from-red-500 via-primary to-red-500 text-transparent bg-clip-text mb-4">
            Welcome to MovSeek!
          </h1>
          <span className="text-lg text-white block mb-6 px-20">
            Whether you are looking for the latest blockbusters or hidden gems, MovSeek helps you find the perfect movie
            to watch. Explore personalized recommendations, search your favorite films, and dive into a world of
            cinematic adventures. Start your movie journey with us today!
          </span>
          <div className="flex justify-center items-center">
            <div className="relative w-full max-w-5xl">
              <input
                type="text"
                placeholder="Search for a movie, tv show, person..."
                className="w-full py-3 pl-5 pr-20 rounded-full shadow-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500 text-lg"
              />
              <button className="absolute right-1 top-1/2 transform -translate-y-1/2 px-5 py-3 bg-gradient-to-r from-cyan-400 to-primary text-white font-geist rounded-full shadow-md transition hover:text-black">
                Search
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto px-10 min-h-screen">
        <main className="flex-grow text-black">
          <section className="container py-6">
            <div className="flex items-center content-center">
              <h1 className="text-4xl text-start font-bold mr-14 mb-6">Trending</h1>
              <div className="relative flex items-center bg-sky-700 p-1 mb-6 rounded-full w-fit">
                <div
                  className="absolute top-1 left-0 right-0 bg-gradient-to-r from-purple-400 to-sky-500 rounded-full border-white border transition-all duration-300 ease-in-out"
                  style={backgroundStyle}
                ></div>
                <button
                  ref={buttonARef}
                  onClick={() => setActiveButton('A')}
                  className={`z-10 px-4 py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
                    activeButton === 'A' ? 'text-white' : 'text-white opacity-60'
                  }`}
                >
                  Today
                </button>
                <button
                  ref={buttonBRef}
                  onClick={() => setActiveButton('B')}
                  className={`z-10 px-4 py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
                    activeButton === 'B' ? 'text-white' : 'text-white opacity-60'
                  }`}
                >
                  This week
                </button>
              </div>
            </div>
            <MovieList category="trending" />
          </section>
          {/* <section className="container py-6">
            <h1 className="text-4xl font-bold mb-6">Xu Hướng</h1>
            <MovieList category="trending" />
          </section> */}
        </main>
      </div>
    </div>
  );
}
