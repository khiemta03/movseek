import Image from 'next/image';

const CastList = ({ cast }: { cast: { name: string; photo: string }[] }) => {
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
      {cast.map((actor, index) => (
        <div key={index} className="flex flex-col items-center">
          <div className="relative w-20 h-20 rounded-full overflow-hidden shadow-md">
            <Image src={actor.photo || '/default-avatar.jpg'} alt={actor.name} layout="fill" objectFit="cover" />
          </div>
          <p className="text-sm font-medium text-center mt-2">{actor.name}</p>
        </div>
      ))}
    </div>
  );
};

export default CastList;
