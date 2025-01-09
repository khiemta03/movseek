package mongo

import (
	"github.com/tmplam/movseek/internal/review"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo implRepository) buildScopeQuery(input review.Scope) bson.M {
	queryFilter := bson.M{}

	if input.UserID != "" {
		queryFilter["user_id"] = input.UserID
	}

	if input.MediaID != 0 {
		queryFilter["media_id"] = input.MediaID
	}

	if input.Type != "" {
		queryFilter["type"] = input.Type
	}

	return queryFilter
}
