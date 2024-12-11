'use client';

import { useSearchParams } from 'next/navigation';
import { useEffect, useState } from 'react';
import {
  fetchSearchMovie,
  // fetchSearchPerson,
  // fetchSearchTV
} from '@/apis/search';
import type {
  SearchMovieResults,
  // SearchPersonResults,
  // SearchTVResults
} from '@/models/search-types';
import MovieSearchCard from '@/components/search/movie-search-card';
import PaginationCustom from '@/components/search/pagination';
import Loading from '@/components/search/search-loading';
// import TVSearchCard from '@/components/search/tv-search-card';
// import PersonSearchCard from '@/components/search/person-search-card';

export default function SearchPage() {
  const searchParams = useSearchParams();
  const query = searchParams.get('query');
  const page = searchParams.get('page');
  const [mode, setMode] = useState<'movie' | 'person' | 'tv'>('movie');
  const [movieResults, setMovieResults] = useState<SearchMovieResults | null>(null);
  // const [tvResults, setTVResults] = useState<SearchTVResults | null>(null);
  // const [personResults, setPersonResults] = useState<SearchPersonResults | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const movieResponse = await fetchSearchMovie(query != null ? query : '', page != null ? parseInt(page) : 1);
        setMovieResults(movieResponse.data);
        // const tvResponse = await fetchSearchTV(newQuery, newPage);
        // setTVResults(tvResponse.data);
        // const personResponse = await fetchSearchPerson(newQuery, newPage);
        // setPersonResults(personResponse.data);
      } catch (err) {
        console.log(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query, page]);

  const handleChangeMode = (newMode: 'movie' | 'person' | 'tv') => {
    if (mode != newMode) {
      setMode(newMode);
    }
  };

  return (
    <div className="flex flex-col min-h-screen font-geist-mono">
      <main className="flex gap-6 container mx-auto mt-5 py-10">
        <div className="w-1/5 text-lg">
          <h2 className="mb-2 font-bold">Search Results For</h2>
          <div className="mb-6">
            <span className="italic mr-5">{`"${query}"`}</span>
          </div>
          <div className="flex flex-col gap-4">
            <div
              className={`p-4 bg-gray-100 ${
                mode == 'movie' ? 'font-bold border border-primary' : ''
              } rounded-lg shadow hover:shadow-lg transition`}
              onClick={() => handleChangeMode('movie')}
            >
              <div className="text-sm flex justify-between">
                <div>Movies</div>
                <div className="font-normal text-gray-500">{movieResults?.total_results || 0}</div>
              </div>
            </div>
            {/* <div
              className={`p-4 bg-gray-100 ${
                mode == 'tv' ? 'font-bold border border-primary' : ''
              } rounded-lg shadow hover:shadow-lg transition`}
              onClick={() => handleChangeMode('tv')}
            >
              <div className="text-sm flex justify-between">
                <div>TV Shows</div>
                <div className="font-normal text-gray-500">{tvResults?.total_results || 0}</div>
              </div>
            </div>
            <div
              className={`p-4 bg-gray-100 ${
                mode == 'person' ? 'font-bold border border-primary' : ''
              } rounded-lg shadow hover:shadow-lg transition`}
              onClick={() => handleChangeMode('person')}
            >
              <div className="text-sm flex justify-between">
                <div>People</div>
                <div className="font-normal text-gray-500">{personResults?.total_results || 0}</div>
              </div>
            </div> */}
          </div>
        </div>
        {movieResults != null && !loading ? (
          <>
            {movieResults.total_results > 0 ? (
              <div className="w-4/5">
                <div className="mb-8 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-4">
                  {movieResults.results.map((movie, index) => (
                    <MovieSearchCard key={index} movie={movie} />
                  ))}
                  {/* {mode == 'movie'
            ? movieResults?.results.map((movie, index) => <MovieSearchCard key={index} movie={movie} />)
            : mode == 'tv'
            ? tvResults?.results.map((tv, index) => <TVSearchCard key={index} tv={tv} />)
            : personResults?.results.map((person, index) => <PersonSearchCard key={index} person={person} />)} */}
                </div>

                {movieResults.total_pages > 1 && (
                  <PaginationCustom
                    currentPage={page != null ? parseInt(page) : 1}
                    totalPage={movieResults.total_pages}
                    query={query != null ? query : ''}
                  />
                )}
              </div>
            ) : (
              <div className="font-bold w-full text-center">There are no movies that matched your query.</div>
            )}
          </>
        ) : (
          <Loading />
        )}
      </main>
    </div>
  );
}
