package mongo

import (
	"context"
	"fmt"

	"github.com/tmplam/movseek/internal/models"
	"github.com/tmplam/movseek/internal/movie/repository"
	"github.com/tmplam/movseek/internal/person"
	"github.com/tmplam/movseek/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	peopleCollection       = "people"
	popularCollection      = "people_popular"
	weekTrendingCollection = "people_trending_week"
	dayTrendingCollection  = "people_trending_day"
)

func (repo implRepository) getPeopleCollection() mongo.Collection {
	return repo.db.Collection(peopleCollection)
}

func (repo implRepository) getTrendingCollection(t string) mongo.Collection {
	return repo.db.Collection(fmt.Sprintf("people_trending_%s", t))
}

func (repo implRepository) getPopularCollection() mongo.Collection {
	return repo.db.Collection(popularCollection)
}

func (repo implRepository) GetOnePerson(ctx context.Context, personID int64) (models.Person, error) {
	col := repo.getPeopleCollection()

	queryFilter := repo.buildPersonQuery(personID)

	var p = models.Person{}
	err := col.FindOne(ctx, queryFilter).Decode(&p)
	if err != nil {
		return models.Person{}, repository.MapError(err)
	}

	return p, err
}

func (repo implRepository) ListPeople(ctx context.Context, input person.ListPeopleOptions) ([]models.Person, error) {
	col := repo.getPeopleCollection()

	queryFilter := repo.buildListPeopleQuery(input)

	findOptions := repo.buildGetPeopleOptions(input.Filter)

	cursor, err := col.Find(ctx, queryFilter, findOptions)
	if err != nil {
		return []models.Person{}, err
	}

	var people []models.Person
	err = cursor.All(ctx, &people)
	if err != nil {
		return []models.Person{}, err
	}

	return people, nil
}

func (repo implRepository) CountPeople(ctx context.Context, input person.ListPeopleOptions) (int, error) {
	col := repo.getPeopleCollection()

	queryFilter := repo.buildListPeopleQuery(input)

	count, err := col.CountDocuments(ctx, queryFilter)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (repo implRepository) GetTrendingPeople(ctx context.Context, input person.GetTrendingPeopleOptions) ([]models.PersonSummary, error) {
	col := repo.getTrendingCollection(input.Type)

	findOptions := repo.buildGetPeopleOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	var people []models.PersonSummary
	err = cursor.All(ctx, &people)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	return people, nil
}

func (repo implRepository) CountTrendingPeople(ctx context.Context, input person.GetTrendingPeopleOptions) (int, error) {
	col := repo.getTrendingCollection(input.Type)

	count, err := col.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (repo implRepository) GetPopularPeople(ctx context.Context, input person.GetPopularPeopleOptions) ([]models.PersonSummary, error) {
	col := repo.getPopularCollection()

	findOptions := repo.buildGetPeopleOptions(input.Filter)

	cursor, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	var people []models.PersonSummary
	err = cursor.All(ctx, &people)
	if err != nil {
		return []models.PersonSummary{}, err
	}

	return people, nil
}

func (repo implRepository) CountPopularPeople(ctx context.Context, input person.GetPopularPeopleOptions) (int, error) {
	col := repo.getPopularCollection()

	count, err := col.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
