'use client';

import { useParams } from 'next/navigation';
import { TMDB_API } from '@/utils/constants';
import { useEffect, useRef, useState } from 'react';
import type { Credits, Keyword, Video } from '@/models/movie-detail-types';
import { handleMovieCredits, selectPreferredVideo } from '@/utils/util-functions/detail-page';
import { Button } from '@/components/ui/button';
import CastList from '@/components/movie/cast-list';
import MainMovieInformationDummy from '@/components/movie/main-movie-information-dummy';
import Trailer from '@/components/movie/trailer';
import { TV } from '@/models/tv-detail-types';
import AltTVInformation from '@/components/tv/alt-tv-information';
import MainTVInformation from '@/components/tv/main-tv-information';
import { fetchTVDetail, fetchTVCredits, fetchTVKeywords, fetchTVVideos } from '@/apis/tv';
import RecommendationList from '@/components/tv/recommendations';
import ReviewsAndRating from '@/components/tv/reviews-and-rating';

export default function TVDetail() {
  const params = useParams();
  const { id } = params;
  const [imageSrc, setImageSrc] = useState('/poster-default.svg');
  const [transitioning, setTransitioning] = useState(false);
  const [transitioningCast, setTransitioningCast] = useState(false);
  const [transitioningRecommendation, setTransitioningRecommendation] = useState(false);
  const [idTV, setIdTV] = useState<number>(0);
  const [tv, setTV] = useState<TV | null>(null);
  const [creadits, setCredits] = useState<Credits | null>(null);
  const [keywords, setKeywords] = useState<Keyword[]>([]);
  const [videos, setVideos] = useState<Video[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [isDisplayFullCastAndCrew, setIsDisplayFullCastAndCrew] = useState<boolean>(false);
  const [viewModeRecommendation, setViewModeRecommendation] = useState<'genres' | 'vectors-search'>('genres');
  const [backgroundStyle, setBackgroundStyle] = useState({});
  const buttonARef = useRef<HTMLButtonElement>(null);
  const buttonBRef = useRef<HTMLButtonElement>(null);
  
  const [isVisible, setIsVisible] = useState(false);

  const toggleVideo = () => setIsVisible(!isVisible);

  useEffect(() => {
    if (id) {
      const numericId = +id;
      setIdTV(numericId);
    }
  }, [id]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const tvResponse = await fetchTVDetail(idTV);
        const data = tvResponse.data.data;
        setImageSrc(TMDB_API.POSTER(data.poster_path));
        setTV(data);

        const creditsResponse = await fetchTVCredits(idTV);
        setCredits(handleMovieCredits(creditsResponse.data));

        const keywordsResponde = await fetchTVKeywords(idTV);
        setKeywords(keywordsResponde.data.results);

        const videoResponse = await fetchTVVideos(idTV);
        setVideos(videoResponse.data.results);
      } catch (err) {
        setError('Failed to fetch movie detail');
        console.log(err);
      } finally {
        setTransitioning(true);
        setTimeout(() => {
          setLoading(false);
          setTransitioning(false);
        }, 500);
      }
    };

    if (idTV > 0) {
      fetchData();
    }
  }, [idTV]);

  useEffect(() => {
    if (isVisible) {
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.overflow = '';
    }
    return () => {
      document.body.style.overflow = '';
    };
  }, [isVisible]);

  const handleModeChangeCast = (mode: boolean) => {
    if (mode !== isDisplayFullCastAndCrew) {
      setTransitioningCast(true);
      setTimeout(() => {
        setIsDisplayFullCastAndCrew(mode);
        setTransitioningCast(false);
      }, 300);
    }
  };

  useEffect(() => {
    const activeRef = viewModeRecommendation === 'genres' ? buttonARef.current : buttonBRef.current;
    if (activeRef) {
      const rect = activeRef.getBoundingClientRect();
      setBackgroundStyle({
        width: `${rect.width}px`,
        transform: `translateX(${activeRef.offsetLeft}px)`,
      });
    }
  }, [viewModeRecommendation]);

  const handleModeChangeRecommendation = (mode: 'genres' | 'vectors-search') => {
    if (mode !== viewModeRecommendation) {
      setViewModeRecommendation(mode);
    }
  };


  if (error)
    return (
      <h1 className="container mx-auto mt-5 font-bold text-2xl">Uh-oh! Something went wrong. Please try later!!!</h1>
    );

  return (
    <div className={`relative transition-opacity duration-500 ${transitioning ? 'opacity-0' : 'opacity-100'}`}>
      {isVisible && (
        <Trailer
          videoId={selectPreferredVideo(videos).key}
          toggleVideo={toggleVideo}
        />
      )}
      {tv !== null && creadits != null && !loading && imageSrc != '/poster-default.svg' ? (
        <div className="font-geist-mono">
          <MainTVInformation
            tv={tv}
            creadits={creadits}
            hasTrailer={videos.length > 0}
            toggleVideo={toggleVideo}
          />

          <div className="flex gap-6 container mx-auto mt-5 py-10">
            <div className={`relative w-4/5`}>
              <div className={`transition-opacity duration-300 ${transitioningCast ? 'opacity-0' : 'opacity-100'}`}>
                {creadits.cast.length >= 6 && (
                  <div className="flex justify-between">
                    {!isDisplayFullCastAndCrew && <h2 className="text-2xl font-bold mb-4">Top Billed Cast</h2>}
                    <div></div>
                    <Button
                      onClick={() => handleModeChangeCast(!isDisplayFullCastAndCrew)}
                      className="text-xl border border-gray-400"
                      variant="outline"
                    >
                      {isDisplayFullCastAndCrew ? 'View less' : 'View Full Cast & Crew'}
                    </Button>
                  </div>
                )}
                <CastList
                  credits={creadits}
                  isfull={isDisplayFullCastAndCrew}
                />
              </div>
              <hr className="my-14 border-t border-gray-300" />
              <div>
                <h2 className="text-2xl font-bold mb-4">Reviews And Rating</h2>
                <ReviewsAndRating
                  rated={null}
                  review={''}
                />
              </div>
              <hr className="my-14 border-t border-gray-300" />
              <div>
                <div className="flex justify-between">
                  <h2 className="text-2xl font-bold mb-4">Recommendations</h2>
                  <div className="relative items-center bg-white border-2 border-blue-500 mb-6 rounded-full w-fit">
                    <div
                      className="absolute h-9 w-[165.33px] z-0 top-0 left-0 bg-blue-500 rounded-full border-white border transition-all duration-300 ease-in-out"
                      style={backgroundStyle}
                    ></div>
                    <button
                      ref={buttonARef}
                      onClick={() => handleModeChangeRecommendation('genres')}
                      className={`relative z-10 px-4 text-sm py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
                        viewModeRecommendation === 'genres' ? 'text-white' : 'text-black opacity-60'
                      }`}
                    >
                      Base on genres
                    </button>
                    <button
                      ref={buttonBRef}
                      onClick={() => handleModeChangeRecommendation('vectors-search')}
                      className={`relative z-10 px-4 text-sm py-2 rounded-full font-medium transition-all duration-300 ease-in-out ${
                        viewModeRecommendation === 'vectors-search' ? 'text-white' : 'text-black opacity-60'
                      }`}
                    >
                      Base on vectors search
                    </button>
                  </div>
                </div>
                <div
                  className={`transition-opacity duration-500 ${
                    transitioningRecommendation ? 'opacity-0' : 'opacity-100'
                  }`}
                >
                  <RecommendationList
                    baseOn={viewModeRecommendation}
                    setTransitioning={setTransitioningRecommendation}
                    tv={tv}
                    keywords={keywords}
                  />
                </div>
              </div>
            </div>
            <div className="w-1/5">
              <AltTVInformation
                tv={tv}
                keywords={keywords}
              />
            </div>
          </div>
        </div>
      ) : (
        <MainMovieInformationDummy />
      )}
    </div>
  );
}
