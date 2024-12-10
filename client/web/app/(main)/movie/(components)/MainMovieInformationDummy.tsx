import Image from 'next/image';
import { TMDB_API } from '@/utils/constants';
import Loading from '../../loading';

const MainMovieInformationDummy = () => {
  return (
    <div className="font-geist-mono">
      <div
        className="relative py-10 px-5 shadow-lg"
        style={{
          backgroundImage: `url(${TMDB_API.POSTER('/poster-default.svg')})`,
          backgroundSize: '100%',
          backgroundRepeat: 'no-repeat',
          backgroundPosition: 'center top',
        }}
      >
        <div className="absolute z-0 inset-0 bg-black/70"></div>
        <div className="relative z-10 container mx-auto text-white">
          <div className="flex flex-row px-20 gap-4">
            <div className="relative w-96 h-[450px] aspect-[2/3] rounded-lg overflow-hidden">
              <Image src="/poster-default.svg" alt={'loading...'} layout="fill" objectFit="contain" />
            </div>
            <div className="flex flex-col justify-center items-center ml-5">
              <Loading />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainMovieInformationDummy;
