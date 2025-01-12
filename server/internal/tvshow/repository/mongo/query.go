package mongo

import (
	"github.com/tmplam/movseek/internal/tvshow"
	"github.com/tmplam/movseek/pkg/date"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo implRepository) buildTVShowQuery(tvID int64) (bson.M, error) {
	queryFilter := bson.M{
		"id": tvID,
	}

	return queryFilter, nil
}

func (repo implRepository) buildFilter(input tvshow.GetTVFilter) bson.M {
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
	queryFilter["first_air_date"] = releaseDateFilter

	if input.GenreIDs != nil {
		queryFilter["genre_ids"] = bson.M{"$in": input.GenreIDs}
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

func (repo implRepository) buildListTVShowsQuery(input tvshow.GetTVFilter, query string) (bson.M, error) {
	queryFilter := repo.buildFilter(input)
	if query != "" {
		queryFilter["$or"] = []bson.M{
			{
				"name": bson.M{
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
			{
				"overview": bson.M{
					"$regex":   query,
					"$options": "i",
				},
			},
			{
				"original_name": bson.M{
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

func (repo implRepository) buildGetTVShowFindOptions(input tvshow.GetTVFilter) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(input.PerPage))
	findOptions.SetSkip(int64((input.Page - 1) * input.PerPage))
	// Create a combined sort specification
	sortSpec := bson.D{}

	if input.TimeOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "first_air_date", Value: input.TimeOrder})
	}
	if input.PopularityOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "popularity", Value: input.PopularityOrder})
	}
	if input.VoteOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "vote_average", Value: input.VoteOrder})
	}
	if input.NameOrder != 0 {
		sortSpec = append(sortSpec, primitive.E{Key: "name", Value: input.NameOrder})
	}

	// Only set sort if we have any sort conditions
	if len(sortSpec) > 0 {
		findOptions.SetSort(sortSpec)
	}
	return findOptions
}
