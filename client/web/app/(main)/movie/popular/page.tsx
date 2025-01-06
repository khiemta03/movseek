'use client';

import { useSearchParams } from 'next/navigation';
import { useEffect, useState } from 'react';
import MovieSearchCard from '@/components/search/movie-search-card';
import Loading from '@/components/search/search-loading';
import { fetchGenresMovie, fetchMoviePopular } from '@/apis/movie-list';
import PaginationCustom from '@/components/person/pagination';
import { FilterSortState, GenresMovieResults, MovieListResults } from '@/models/movie-list-types';
import { ChevronDown } from 'lucide-react';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import DoubleSlider from '@/components/movie-list/double-sider';
import { deepEqual } from '@/utils/util-functions/movie-list-page';
// import { format, parse } from 'date-fns';
import { Input } from '@/components/ui/input';

export default function SearchPage() {
  const searchParams = useSearchParams();
  const page = searchParams.get('page');
  const [movieResults, setMovieResults] = useState<MovieListResults | null>(null);
  const [genreListResults, setGenreListResults] = useState<GenresMovieResults | null>(null);
  const [loading, setLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [isOpenSort, setIsOpenSort] = useState(false);
  const [isOpenFilter, setIsOpenFilter] = useState(false);
  const initFilterSortState = {
    sort: 'popularity-desc',
    genre: [],
    releaseDate: {
      from: null,
      to: new Date(),
    },
    userScore: {
      from: 0,
      to: 100,
    },
    runTime: {
      from: 0,
      to: 360,
    },
  };
  const [filterSortState, setFilterSortState] = useState<FilterSortState>(initFilterSortState);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const movieResponse = await fetchMoviePopular(page != null ? parseInt(page) : 1);
        console.log(movieResponse.data);
        setMovieResults(movieResponse.data);
        const genreResponse = await fetchGenresMovie();
        console.log(genreResponse.data);
        setGenreListResults(genreResponse.data);
      } catch (err) {
        console.log(err);
        setIsError(true);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [page]);

  useEffect(() => {
    console.log(filterSortState);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filterSortState]);

  const updateFilterSort = (key: keyof FilterSortState, value: unknown) => {
    setFilterSortState((prev) => ({
      ...prev,
      [key]: value,
    }));
  };

  // Cập nhật thuộc tính lồng nhau (nested)
  const updateNestedFilterSort = <T extends keyof FilterSortState>(
    parentKey: T,
    childKey: keyof FilterSortState[T],
    value: unknown,
  ) => {
    if (typeof filterSortState[parentKey] === 'object' && filterSortState[parentKey] !== null) {
      setFilterSortState((prev) => ({
        ...prev,
        [parentKey]: {
          ...(prev[parentKey] as Record<string, unknown>),
          [childKey]: value,
        },
      }));
    } else {
      console.error(`${parentKey} is not an object`);
    }
  };

  const handleGenreClick = (genreId: number) => {
    setFilterSortState((prevState) => {
      const newGenres = prevState.genre.includes(genreId)
        ? prevState.genre.filter((id) => id !== genreId)
        : [...prevState.genre, genreId];

      return {
        ...prevState,
        genre: newGenres,
      };
    });
  };

  const handleUserScoreChange = (newUserScore: number[]) => {
    updateNestedFilterSort('userScore', 'from', newUserScore[0]);
    updateNestedFilterSort('userScore', 'to', newUserScore[1]);
  };

  const handleRuntimeChange = (newRuntime: number[]) => {
    updateNestedFilterSort('runTime', 'from', newRuntime[0]);
    updateNestedFilterSort('runTime', 'to', newRuntime[1]);
  };

  const handleDateChange = (e: React.ChangeEvent<HTMLInputElement>, key: 'from' | 'to') => {
    const dateValue = new Date(e.target.value);
    if (!isNaN(dateValue.getTime())) {
      if (key === 'from' && filterSortState.releaseDate.to && dateValue > filterSortState.releaseDate.to) {
        alert('The "from" date cannot be greater than the "to" date.');
        return;
      }
      if (key === 'to' && filterSortState.releaseDate.from && dateValue < filterSortState.releaseDate.from) {
        alert('The "to" date cannot be earlier than the "from" date.');
        return;
      }
      updateNestedFilterSort('releaseDate', key, dateValue);
    } else {
      updateNestedFilterSort('releaseDate', key, null);
    }
  };

  if (isError)
    return (
      <h1 className="container mx-auto mt-5 font-bold text-2xl">Uh-oh! Something went wrong. Please try later!!!</h1>
    );

  return (
    <div className="flex flex-col min-h-screen font-geist-mono">
      <main className="container mx-auto py-10">
        <h2 className="mb-5 text-2xl font-bold ">Popular Movies</h2>
        <div className="flex gap-6">
          <div className="w-1/5 text-2xl">
            <div className="flex flex-col gap-8">
              <div className="flex flex-col">
                <div className={`p-4 bg-gray-100 rounded-lg shadow hover:shadow-lg transition border`}>
                  <Collapsible open={isOpenSort} onOpenChange={setIsOpenSort} className="space-y-2">
                    <CollapsibleTrigger asChild>
                      <div className="flex justify-between hover:cursor-pointer">
                        <div className="text-lg font-bold">Sort</div>
                        <div className={`transform transition-transform ${isOpenSort ? 'rotate-0' : '-rotate-90'}`}>
                          <ChevronDown />
                        </div>
                      </div>
                    </CollapsibleTrigger>
                    <CollapsibleContent className="space-y-2">
                      <hr className="my-2 border-t border-gray-300" />
                      <div className="text-base">Sort Results By</div>
                      <Select value={filterSortState.sort} onValueChange={(value) => updateFilterSort('sort', value)}>
                        <SelectTrigger className="w-full text-xs bg-white">
                          <SelectValue placeholder="Select a fruit" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            <SelectItem value="popularity-desc">Popularity Descending</SelectItem>
                            <SelectItem value="popularity-asc">Popularity Ascending</SelectItem>
                            <SelectItem value="rating-desc">Rating Descending</SelectItem>
                            <SelectItem value="rating-asc">Rating Ascending</SelectItem>
                            <SelectItem value="release-date-desc">Release Date Descending</SelectItem>
                            <SelectItem value="release-date-asc">Release Date Ascending</SelectItem>
                            <SelectItem value="title-a-z">Title (A-Z)</SelectItem>
                            <SelectItem value="title-z-a">Title (Z-A)</SelectItem>
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    </CollapsibleContent>
                  </Collapsible>
                </div>
              </div>
              <div className="flex flex-col">
                <div className={`p-4 bg-gray-100 rounded-lg shadow hover:shadow-lg transition`}>
                  <Collapsible open={isOpenFilter} onOpenChange={setIsOpenFilter} className="space-y-2">
                    <CollapsibleTrigger asChild>
                      <div className="flex justify-between hover:cursor-pointer">
                        <div className="text-lg font-bold">Filters</div>
                        <div className={`transform transition-transform ${isOpenFilter ? 'rotate-0' : '-rotate-90'}`}>
                          <ChevronDown />
                        </div>
                      </div>
                    </CollapsibleTrigger>
                    <CollapsibleContent className="space-y-2">
                      <div>
                        <hr className="my-2 border-t border-gray-300" />
                        <div className="text-base mb-3">Release Dates</div>
                        <div className="flex justify-between items-center mb-3">
                          <div className="text-sm">From</div>
                          <Input
                            type="date"
                            value={
                              filterSortState.releaseDate.from
                                ? filterSortState.releaseDate.from.toISOString().split('T')[0]
                                : ''
                            }
                            onChange={(e) => handleDateChange(e, 'from')}
                            placeholder="YYYY-MM-DD"
                            className="w-40 bg-white"
                          />
                        </div>
                        <div className="flex justify-between items-center mb-3">
                          <div className="text-sm">To</div>
                          <Input
                            type="date"
                            value={
                              filterSortState.releaseDate.to
                                ? filterSortState.releaseDate.to.toISOString().split('T')[0]
                                : ''
                            }
                            onChange={(e) => handleDateChange(e, 'to')}
                            placeholder="YYYY-MM-DD"
                            className="w-40 bg-white"
                          />
                        </div>
                      </div>
                      <div>
                        <hr className="my-2 border-t border-gray-300" />
                        <div className="text-base">Genres</div>
                        {genreListResults != null && !loading && (
                          <div className="flex flex-wrap items-center">
                            {genreListResults.genres.map((genre, index) => (
                              <Button
                                key={index}
                                variant={'outline'}
                                size={'sm'}
                                className={`mr-2 mt-2 border text-xs rounded-full ${
                                  filterSortState.genre.includes(genre.id) &&
                                  'bg-primary text-white hover:bg-primary-dark hover:text-white'
                                } border-black`}
                                onClick={() => handleGenreClick(genre.id)}
                              >
                                {genre.name}
                              </Button>
                            ))}
                          </div>
                        )}
                      </div>
                      <div>
                        <hr className="my-2 border-t border-gray-300" />
                        <div className="text-base mb-3">User Score</div>
                        <DoubleSlider
                          values={[filterSortState.userScore.from, filterSortState.userScore.to]}
                          step={1}
                          min={0}
                          max={100}
                          onChange={handleUserScoreChange}
                        />
                      </div>
                      <div>
                        <hr className="my-2 border-t border-gray-300" />
                        <div className="text-base mb-3">Runtime (minutes)</div>
                        <DoubleSlider
                          values={[filterSortState.runTime.from, filterSortState.runTime.to]}
                          step={1}
                          min={0}
                          max={360}
                          onChange={handleRuntimeChange}
                        />
                      </div>
                    </CollapsibleContent>
                  </Collapsible>
                </div>
              </div>
              <Button disabled={deepEqual(initFilterSortState, filterSortState)}>Search</Button>
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
                  </div>
                  {movieResults.total_pages > 1 && (
                    <PaginationCustom
                      currentPage={page != null ? parseInt(page) : 1}
                      totalPage={movieResults.total_pages}
                      endpoint={'/movie/popular'}
                    />
                  )}
                </div>
              ) : (
                <div className="font-bold w-full text-center">There are no movies to display.</div>
              )}
            </>
          ) : (
            <Loading />
          )}
        </div>
      </main>
    </div>
  );
}
