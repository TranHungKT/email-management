package subscriberHandlers

import (
	"context"
	"strings"

	"github.com/TranHungKT/email_management/database"
	"github.com/TranHungKT/email_management/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewSubscriberHandler(newSubscriber models.NewSubscriberRequestPayload, lists []models.List) (primitive.ObjectID, error) {
	if newSubscriber.Status == "" {
		newSubscriber.Status = models.SubscriberStatusEnabled
	}

	var subscribedLists = make([]models.SubscribedList, 0)

	for _, list := range lists {
		subStatus := models.SubscriptionStatusConfirmed

		if list.Optin == models.ListOptinDouble {
			subStatus = models.SubscriptionStatusUnConfirmed
		}

		subscribedList := models.SubscribedList{
			ListId:             list.Id,
			SubscriptionStatus: subStatus,
		}
		subscribedLists = append(subscribedLists, subscribedList)
	}

	newSubscriber.Name = strings.TrimSpace(newSubscriber.Name)

	var subscriber = models.Subscriber{
		Email:      newSubscriber.Email,
		Name:       newSubscriber.Name,
		Attributes: newSubscriber.Attributes,
		Lists:      subscribedLists,
		Status:     newSubscriber.Status,
	}

	result, err := database.SubscriberCollection().InsertOne(context.TODO(), &subscriber)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
