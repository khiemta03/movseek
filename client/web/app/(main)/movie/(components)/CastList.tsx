import { Credits } from '@/utils/types';
import CastCard from './CastCard';

interface CastListProps {
  credits: Credits;
}

const CastList: React.FC<CastListProps> = ({ credits }) => {
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
      {credits.cast.map((actor, index) => (
        <CastCard key={index} actor={actor} />
      ))}
    </div>
  );
};

export default CastList;
