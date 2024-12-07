import MovieList from '@/components/MovieList';

export default function Home() {
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

      <div className="container mx-auto flex flex-col min-h-screen">
        <main className="flex-grow text-black">
          <section className="container mx-auto py-6">
            <h1 className="text-4xl text-start font-bold mb-6">Trending</h1>
            <MovieList category="popular" />
          </section>
          <section className="container mx-auto py-6">
            <h1 className="text-4xl font-bold mb-6">Xu Hướng</h1>
            <MovieList category="trending" />
          </section>
        </main>
      </div>
    </div>
  );
}
