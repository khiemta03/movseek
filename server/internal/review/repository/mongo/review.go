package mongo

import (
	"context"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/review"
	"github.com/tmplam/movseek/pkg/mongo"
)

const (
	commentsCollection = "comments"
	ratingsCollection  = "ratings"
)

func (repo implRepository) getCommentsCollection() mongo.Collection {
	return repo.db.Collection(commentsCollection)
}

func (repo implRepository) getRatingsCollection() mongo.Collection {
	return repo.db.Collection(ratingsCollection)
}

func (repo implRepository) AddComment(ctx context.Context, input review.AddCommentInput) (models.Comment, error) {
	col := repo.getCommentsCollection()

	model := repo.buildAddComment(input)

	_, err := col.InsertOne(ctx, model)
	if err != nil {
		return models.Comment{}, err
	}

	return model, nil
}

func (repo implRepository) UpdateComment(ctx context.Context, input review.UpdateCommentInput) error {
	col := repo.getCommentsCollection()

	queryFilter := repo.buildScopeQuery(input.Scope)
	update := repo.buildUpdateComment(input)

	_, err := col.UpdateOne(ctx, queryFilter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo implRepository) DeleteComment(ctx context.Context, input review.Scope) error {
	col := repo.getCommentsCollection()

	queryFilter := repo.buildScopeQuery(input)

	_, err := col.DeleteOne(ctx, queryFilter)
	if err != nil {
		return err
	}

	return nil
}

func (repo implRepository) AddRating(ctx context.Context, input review.AddRatingInput) (models.Rating, error) {
	col := repo.getRatingsCollection()

	model := repo.buildAddRating(input)

	_, err := col.InsertOne(ctx, model)
	if err != nil {
		return models.Rating{}, err
	}

	return model, nil
}

func (repo implRepository) UpdateRating(ctx context.Context, input review.UpdateRatingInput) error {
	col := repo.getRatingsCollection()

	queryFilter := repo.buildScopeQuery(input.Scope)
	update := repo.buildUpdateRating(input)

	_, err := col.UpdateOne(ctx, queryFilter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo implRepository) DeleteRating(ctx context.Context, input review.Scope) error {
	col := repo.getRatingsCollection()

	queryFilter := repo.buildScopeQuery(input)

	_, err := col.DeleteOne(ctx, queryFilter)
	if err != nil {
		return err
	}

	return nil
}

func (repo implRepository) GetComments(ctx context.Context, input review.Scope) ([]models.Comment, error) {
	col := repo.getCommentsCollection()

	queryFilter := repo.buildScopeQuery(input)

	cursor, err := col.Find(ctx, queryFilter)
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (repo implRepository) GetRatings(ctx context.Context, input review.Scope) ([]models.Rating, error) {
	col := repo.getRatingsCollection()

	queryFilter := repo.buildScopeQuery(input)

	cursor, err := col.Find(ctx, queryFilter)
	if err != nil {
		return nil, err
	}

	var ratings []models.Rating
	err = cursor.All(ctx, &ratings)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
