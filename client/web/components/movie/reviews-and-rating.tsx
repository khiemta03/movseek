import { useEffect, useState } from 'react';
import { Slider } from '@/components/ui/slider';
import { CircleMinus, Star } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import { useToast } from '@/hooks/use-toast';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { Movie } from '@/models/movie-detail-types';
import { ToastAction } from '@/components/ui/toast';
import { useRouter } from 'next/navigation';
import { usePathname } from 'next/navigation';
import { addComment, deleteComment, getCommentsByMedia, updateComment } from '@/apis/reviews';
import { Review } from '@/models/review-types';

interface ReviewsAndRatingProps {
  movie: Movie;
  isSignedIn: boolean;
  user_id: string;
  avatar: string;
  username: string;
}

const ReviewsAndRating: React.FC<ReviewsAndRatingProps> = ({ movie, isSignedIn, user_id, avatar, username }) => {
  const router = useRouter();
  const pathname = usePathname();
  const { toast } = useToast();
  const [myRating, setMyRating] = useState<number | null>(null);
  const [isRated, setIsRated] = useState<boolean>(false);
  const [myComment, setMyComment] = useState<string>('');
  const [isCommented, setIsCommented] = useState<boolean>(false);
  const [reviews, setReviews] = useState<Review[]>([]);
  const [ratings, setRatings] = useState([]);

  const handleCommentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const value = e.target.value;
    if (value.length <= 2000) {
      setMyComment(value);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const commentsResponse = await getCommentsByMedia(parseInt(movie.id), 'movie');
        console.log('>>>', commentsResponse.data.data.comments);
        if (Array.isArray(commentsResponse.data.data.comments)) {
          setReviews(commentsResponse.data.data.comments);
          const comments: Review[] = commentsResponse.data.data.comments;
          const comment = comments.find((comment) => comment.user_id === user_id);
          console.log(comment);
          // if (index >= 0) {
          //   setIsCommented(true);
          //   setMyComment(comments[index].comment);
          // }
        }
        // const watchlistItemResponse = await getWatchlistItem(user_id);
        // if (Array.isArray(watchlistItemResponse.data.data.movie_id)) {
        //   if (watchlistItemResponse.data.data.movie_id.includes(parseInt(movie.id))) {
        //     setWatchlist(true);
        //   } else {
        //     setWatchlist(false);
        //   }
        // }
      } catch (err) {
        console.log(err);
      }
    };

    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // const handleSubmit = async (e: React.FormEvent) => {
  //   e.preventDefault();
  //   if (comment.trim() === '' && rating == null) return;
  //   console.log(isSignedIn);
  //   if (isSignedIn) {
  //     if (rating) {
  //       //   await addSavedItem(parseInt(movie.id), 'movie', 'favorite', user_id);
  //       // } else {
  //       //   await removeSavedItem(parseInt(movie.id), 'movie', 'favorite', user_id);
  //       // }
  //       // setFavorite(!favorite);
  //       // toast({
  //       //   title: 'Success',
  //       //   description: `${movie.title} was ${favorite ? 'removed from' : 'added to'} your favourite list.`,
  //       //   duration: 3000,
  //       //   className: 'bg-green-600 text-white border border-gray-200',
  //       // });
  //       if (comment.trim() != '') {
  //         // const response = await addComment(avatar, comment, parseInt(movie.id), 'movie', user_id, username);
  //         console.log(avatar, comment, parseInt(movie.id), 'movie', user_id, username);
  //       }
  //       if (rating != null) {
  //       }

  //       setComment('');
  //     }
  //   } else {
  //     toast({
  //       title: 'Login Required',
  //       description: 'You must log in to perform this action.',
  //       duration: 3000,
  //       className: 'bg-primary text-white border border-gray-200',
  //       action: (
  //         <ToastAction
  //           altText="Go to login"
  //           className="text-primary bg-white hover:bg-gray-200 px-3 py-1 rounded-lg"
  //           onClick={() => {
  //             router.push(`/sign-in?redirect=${encodeURIComponent(pathname)}`);
  //           }}
  //         >
  //           Sign in
  //         </ToastAction>
  //       ),
  //     });
  //   }
  // };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // if (myComment.trim() === '' && myRating == null) return;
    if (isSignedIn) {
      try {
        if (myComment.trim() !== '') {
          if (isCommented) {
            const response = await updateComment(avatar, myComment, parseInt(movie.id), 'movie', user_id, username);
            console.log(response);
            toast({
              title: 'Success',
              description: 'Your comment has been updated.',
              duration: 3000,
              className: 'bg-green-600 text-white border border-gray-200',
            });
          } else {
            const response = await addComment(avatar, myComment, parseInt(movie.id), 'movie', user_id, username);
            console.log(response);
            toast({
              title: 'Success',
              description: 'Your comment has been added.',
              duration: 3000,
              className: 'bg-green-600 text-white border border-gray-200',
            });
          }
        } else if (isCommented) {
          await deleteComment(user_id, parseInt(movie.id), 'movie');
          toast({
            title: 'Success',
            description: 'Your comment has been deleted.',
            duration: 3000,
            className: 'bg-red-600 text-white border border-gray-200',
          });
        }

        // Xử lý rating
        if (myRating != null) {
          if (isRated) {
            // Nếu đã có rating, sửa rating
            // await updateRating(existingRating.id, rating, user_id);
            // toast({
            //   title: 'Success',
            //   description: 'Your rating has been updated.',
            //   duration: 3000,
            //   className: 'bg-green-600 text-white border border-gray-200',
            // });
          } else {
            // Nếu chưa có rating, thêm rating mới
            // await addRating(rating, parseInt(movie.id), 'movie', user_id);
            // toast({
            //   title: 'Success',
            //   description: 'Your rating has been added.',
            //   duration: 3000,
            //   className: 'bg-green-600 text-white border border-gray-200',
            // });
          }
        } else if (isRated) {
          // Nếu không có rating mới và đã có rating trước đó, xóa rating
          // await deleteRating(existingRating.id, user_id);
          // toast({
          //   title: 'Success',
          //   description: 'Your rating has been deleted.',
          //   duration: 3000,
          //   className: 'bg-red-600 text-white border border-gray-200',
          // });
        }
        // Reset trạng thái comment và rating
        // setMyComment('');
        // setMyRating(null);
      } catch (error) {
        toast({
          title: 'Error',
          description: 'An error occurred. Please try again later.',
          duration: 3000,
          className: 'bg-red-600 text-white border border-gray-200',
        });
        console.error(error);
      }
    } else {
      toast({
        title: 'Login Required',
        description: 'You must log in to perform this action.',
        duration: 3000,
        className: 'bg-primary text-white border border-gray-200',
        action: (
          <ToastAction
            altText="Go to login"
            className="text-primary bg-white hover:bg-gray-200 px-3 py-1 rounded-lg"
            onClick={() => {
              router.push(`/sign-in?redirect=${encodeURIComponent(pathname)}`);
            }}
          >
            Sign in
          </ToastAction>
        ),
      });
    }
  };

  return (
    <div>
      <div className="my-2">
        <form
          onSubmit={handleSubmit}
          className="w-full mx-auto p-4 border rounded-lg shadow-lg bg-white space-y-6"
        >
          <div>
            <div className="space-y-2">
              <div className="flex justify-between items-center">
                <h2 className="text-lg font-semibold">{`${isCommented ? 'My review' : 'Write a review'}`}</h2>
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Button
                        variant={'ghost'}
                        size={'icon'}
                        onClick={() => setMyComment('')}
                      >
                        <CircleMinus />
                      </Button>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>Clear my review</p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </div>
              <Textarea
                placeholder="Write your review here..."
                value={myComment}
                rows={3}
                onChange={handleCommentChange}
                maxLength={2000}
                className="w-full"
              />
              <div className="text-right text-sm text-gray-500">{myComment.length}/2000</div>
            </div>
            <div className="space-y-2">
              <h2 className="text-lg font-semibold">Rating</h2>
              <div className="w-full flex items-center gap-5">
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Button
                        variant={'ghost'}
                        size={'icon'}
                        onClick={() => setMyRating(null)}
                      >
                        <CircleMinus />
                      </Button>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>Clear my vote</p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
                <Slider
                  value={[myRating != null ? myRating : 0]}
                  max={100}
                  step={10}
                  onValueChange={(value) => setMyRating(value[0])}
                />
                <Button
                  variant="default"
                  size="icon"
                  className={`rounded-full border border-gray-200 p-6 mr-4`}
                  disabled={myRating == null}
                >
                  {myRating != null ? myRating : <Star className="font-bold fill-white" />}
                </Button>
              </div>
            </div>
          </div>
          <Button
            type="submit"
            disabled={myComment.trim() === '' && myRating == null}
            className="w-full"
          >
            Save
          </Button>
        </form>
      </div>
    </div>
  );
};

export default ReviewsAndRating;
