package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromHexOrNil(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func ObjectIDsFromHexOrNil(ids []string) []primitive.ObjectID {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIDs[i] = ObjectIDFromHexOrNil(id)
	}
	return objIDs
}
