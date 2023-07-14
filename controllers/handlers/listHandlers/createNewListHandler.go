package listHandlers

import (
	"context"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewListHandler(list models.List) (primitive.ObjectID, error) {
	if list.Type == "" {
		list.Type = models.ListTypePrivate
	}

	if list.Optin == "" {
		list.Optin = models.ListOptinSingle
	}

	cursor, err := database.ListCollection().InsertOne(context.TODO(), &list)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	return cursor.InsertedID.(primitive.ObjectID), nil
}
