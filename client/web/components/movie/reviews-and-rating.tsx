import { useState } from 'react';
import { Slider } from '@/components/ui/slider';
import { CircleMinus, Star } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import { useToast } from '@/hooks/use-toast';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';

interface ReviewsAndRatingProps {
  rated: number | null;
  review: string;
}

const ReviewsAndRating: React.FC<ReviewsAndRatingProps> = ({ rated, review }) => {
  const { toast } = useToast();
  const [rating, setRating] = useState<number | null>(rated);
  const [comment, setComment] = useState(review);

  const handlClearMyVote = (rating: number | null) => {
    if (rating == null) {
      toast({
        title: 'Warning',
        description: `You have no reviews yet.`,
        duration: 3000,
        className: 'bg-yellow-500 text-white border border-gray-200',
      });
      return;
    }
    //call api here
    setRating(null);
    toast({
      title: 'Success',
      description: `Your rating has been cleared.`,
      duration: 3000,
      className: 'bg-green-600 text-white border border-gray-200',
    });
  };

  // const handlVote = (rating: number | null) => {
  //   if (rating == null) {
  //     toast({
  //       title: 'Warning',
  //       description: `Please rate before voting.`,
  //       duration: 3000,
  //       className: 'bg-yellow-500 text-white border border-gray-200',
  //     });
  //     return;
  //   }
  //   //call api here
  //   setRating(rating);
  //   toast({
  //     title: 'Success',
  //     description: `Your rating has been saved.`,
  //     duration: 3000,
  //     className: 'bg-green-600 text-white border border-gray-200',
  //   });
  // };

  const handleCommentChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const value = e.target.value;
    if (value.length <= 2000) {
      setComment(value);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (comment.trim() === '' && rating == null) return;

    console.log('Submitted Comment:', comment);
    setComment(''); // Reset input sau khi gửi
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
              <h2 className="text-lg font-semibold">Write a review</h2>
              <Textarea
                placeholder="Write your review here..."
                value={comment}
                rows={3}
                onChange={handleCommentChange}
                maxLength={2000}
                className="w-full"
              />
              <div className="text-right text-sm text-gray-500">{comment.length}/2000</div>
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
                        onClick={() => handlClearMyVote(rating)}
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
                  value={[rating != null ? rating : 0]}
                  max={100}
                  step={10}
                  onValueChange={(value) => setRating(value[0])}
                />
                <Button
                  variant="default"
                  size="icon"
                  className={`rounded-full border border-gray-200 p-6 mr-4`}
                  disabled={rating == null}
                >
                  {rating != null ? rating : <Star className="font-bold fill-white" />}
                </Button>
              </div>
            </div>
          </div>
          <Button
            type="submit"
            disabled={comment.trim() === '' && rating == null}
            className="w-full"
          >
            Submit
          </Button>
        </form>
      </div>
    </div>
  );
};

export default ReviewsAndRating;
