package listHandlers

import (
	"context"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetListByIdsHandler(listIds []primitive.ObjectID) ([]models.List, error) {
	cursor, err := database.ListCollection().Find(context.TODO(), bson.D{primitive.E{Key: "_id", Value: bson.D{primitive.E{Key: "$in", Value: listIds}}}})

	if err != nil {
		return make([]models.List, 0), err
	}

	var lists []models.List

	if err = cursor.All(context.TODO(), &lists); err != nil {
		return lists, err
	}

	return lists, nil
}
