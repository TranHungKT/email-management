package listHandlers

import (
	"context"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
)

func CreateNewListHandler(list models.List) (interface{}, error) {
	if list.Type == "" {
		list.Type = models.ListTypePrivate
	}

	if list.Optin == "" {
		list.Optin = models.ListOptinSingle
	}

	cursor, err := database.ListCollection().InsertOne(context.TODO(), &list)

	if err != nil {
		return cursor.InsertedID, err
	}

	return cursor.InsertedID, nil
}
