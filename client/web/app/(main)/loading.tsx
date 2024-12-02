import { Skeleton } from '@/components/ui/skeleton';

const Loading = () => {
  return (
    <div className="flex justify-center mt-6">
      <div className="flex flex-col space-y-3">
        <Skeleton className="h-[125px] w-[450px] max-w-full rounded-xl" />
        <div className="space-y-2">
          <Skeleton className="h-4 w-[450px] max-w-full" />
          <Skeleton className="h-4 w-[300px] max-w-full" />
        </div>
      </div>
    </div>
  );
};

export default Loading;
