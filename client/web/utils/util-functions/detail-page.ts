/* eslint-disable @typescript-eslint/no-explicit-any */
import { Crew, CrewGroupedByDepartment, Video } from '@/models/movie-detail-types';

function convertMinutes(minutes: number): string {
  const hours = Math.floor(minutes / 60);
  const minutesRemaining = minutes % 60;
  return `${hours}h ${minutesRemaining}m`;
}

function getCrewByJob(crewGrouped: CrewGroupedByDepartment, department: string, job: string): Crew[] {
  const departmentCrew = crewGrouped[department];
  if (!departmentCrew) return [];

  return departmentCrew.filter((member) => member.job === job);
}

function pickMovieFields(movie: any) {
  return {
    backdrop_path: movie.backdrop_path,
    budget: movie.budget,
    genres: movie.genres,
    id: movie.id,
    homepage: movie.homepage,
    origin_country: movie.origin_country,
    original_language: movie.original_language,
    original_title: movie.original_title,
    overview: movie.overview,
    popularity: movie.popularity,
    poster_path: movie.poster_path,
    release_date: movie.release_date,
    revenue: movie.revenue,
    runtime: movie.runtime,
    status: movie.status,
    tagline: movie.tagline,
    title: movie.title,
    vote_average: movie.vote_average,
    vote_count: movie.vote_count,
  };
}

function handleMovieCredits(data: any) {
  const cast = data.cast.map((member: any) => ({
    id: member.id,
    name: member.name,
    gender: member.gender,
    original_name: member.original_name,
    popularity: member.popularity,
    profile_path: member.profile_path,
    cast_id: member.cast_id,
    character: member.character,
    credit_id: member.credit_id,
    order: member.order,
  }));

  const crew = data.crew.reduce((grouped: any, member: any) => {
    if (!grouped[member.department]) {
      grouped[member.department] = [];
    }
    grouped[member.department].push({
      id: member.id,
      name: member.name,
      gender: member.gender,
      original_name: member.original_name,
      popularity: member.popularity,
      profile_path: member.profile_path,
      credit_id: member.credit_id,
      department: member.department,
      job: member.job,
    });
    return grouped;
  }, {});

  return { cast, crew };
}

function formatCurrency(amount: number): string {
  return `$${amount.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
}

function getSizeOfCrew(crewGrouped: CrewGroupedByDepartment): number {
  return Object.values(crewGrouped).reduce((total, crewArray) => total + crewArray.length, 0);
}

function selectPreferredVideo(results: Video[]) {
  const youtubeVideos = results.filter((video) => video.site === 'YouTube');

  const trailerVideo = youtubeVideos.find((video) => video.type === 'Trailer');
  if (trailerVideo) return trailerVideo;

  const teaserVideo = youtubeVideos.find((video) => video.type === 'Teaser');
  if (teaserVideo) return teaserVideo;

  return results[0];
}

export {
  getSizeOfCrew,
  formatCurrency,
  convertMinutes,
  getCrewByJob,
  pickMovieFields,
  handleMovieCredits,
  selectPreferredVideo,
};
