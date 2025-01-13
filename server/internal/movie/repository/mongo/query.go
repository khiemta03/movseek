package mongo

import (
	"github.com/tmplam/movseek/internal/movie"
	"github.com/tmplam/movseek/pkg/date"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildMovieQuery(movieID int64) (bson.M, error) {
	queryFilter := bson.M{
		"id": movieID,
	}

	return queryFilter, nil
}

func (repo implRepository) buildFilter(input movie.GetMovieFilter) bson.M {
	queryFilter := bson.M{}

	releaseDateFilter := bson.M{
		"$lte": date.GetCurrentDate(),
	}
	if input.StartDate != "" {
		releaseDateFilter["$gte"] = input.StartDate
	}
	if input.EndDate != "" {
		releaseDateFilter["$lte"] = input.EndDate
	}
	queryFilter["release_date"] = releaseDateFilter

	if input.GenreIDs != nil {
		queryFilter["$or"] = []bson.M{
			{
				"genres": bson.M{
					"$elemMatch": bson.M{
						"id": bson.M{"$in": input.GenreIDs},
					},
				},
			},
			{
				"genre_ids": bson.M{
					"$in": input.GenreIDs,
				},
			},
		}
	}

	voteAverageFilter := bson.M{}

	if input.StartAverageVote > 0 {
		voteAverageFilter["$gte"] = input.StartAverageVote
	}

	if input.EndAverageVote > 0 {
		voteAverageFilter["$lte"] = input.EndAverageVote
	}

	if len(voteAverageFilter) > 0 {
		queryFilter["vote_average"] = voteAverageFilter
	}

	return queryFilter
}

func (repo implRepository) buildListMoviesQuery(input movie.GetMovieFilter, query string) (bson.M, error) {
	queryFilter := repo.buildFilter(input)
	if query != "" {
		queryFilter["$or"] = []bson.M{
			{
				"title": bson.M{
					"$regex":   query,
					"$options": "i",
				},
			},
			{
				"keywords.name": bson.M{
					"$regex":   query,
					"$options": "i",
				},
			},
			{
				"overview": bson.M{
					"$regex":   query,
					"$options": "i",
				},
			},
			{
				"belongs_to_collection.name": bson.M{
					"$regex":   query,
					"$options": "i",
				},
			},
		}
	}

	if len(input.IDs) > 0 {
		queryFilter["id"] = bson.M{"$in": input.IDs}
	}

	if len(input.ObjectIDs) > 0 {
		objectIDs := make([]primitive.ObjectID, len(input.ObjectIDs))
		for i, id := range input.ObjectIDs {
			objectID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return bson.M{}, err
			}
			objectIDs[i] = objectID
		}
		queryFilter["_id"] = bson.M{"$in": objectIDs}
	}

	return queryFilter, nil
}

func (repo implRepository) buildGetMovieFindOptions(input movie.GetMovieFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))

	// Create a combined sort specification
	sortSpec := bson.D{}

	if input.TimeOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "release_date", Value: input.TimeOrder})
	}
	if input.PopularityOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "popularity", Value: input.PopularityOrder})
	}
	if input.VoteOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "vote_average", Value: input.VoteOrder})
	}
	if input.TitleOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "title", Value: input.TitleOrder})
	}

	// Only set sort if we have any sort conditions
	if len(sortSpec) > 0 {
		findOptions.SetSort(sortSpec)
	}

	return findOptions
}

func (repo implRepository) buildGetMovieGenresQuery(objectIDs []string) (bson.M, error) {
	queryFilter := bson.M{}

	if len(objectIDs) > 0 {
		res := make([]primitive.ObjectID, len(objectIDs))
		for i, id := range objectIDs {
			objectID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return bson.M{}, err
			}
			res[i] = objectID
		}
		queryFilter["_id"] = bson.M{"$in": res}
	}

	return queryFilter, nil
}
