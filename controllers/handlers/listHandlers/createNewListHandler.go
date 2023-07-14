package listHandlers

import (
	"context"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
)

func CreateNewListHandler(list models.List) interface{} {
	if list.Type == "" {
		list.Type = models.ListTypePrivate
	}

	if list.Optin == "" {
		list.Optin = models.ListOptinSingle
	}

	cursor, err := database.ListCollection().InsertOne(context.TODO(), &list)

	if err != nil {
		panic(err)
	}

	return cursor.InsertedID
}
