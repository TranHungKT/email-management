package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SubscriberStatusEnabled     = "enabled"
	SubscriberStatusDisabled    = "disabled"
	SubscriberStatusBlockListed = "blockListed"
)

type Subscriber struct {
	Base       `bson:",inline"`
	Email      string `validate:"required,email"`
	Name       string `validate:"required,max=200"`
	Attributes map[string]interface{}
	Status     string
	Lists      []primitive.ObjectID
}

func (subscriber *Subscriber) MarshalBSON() ([]byte, error) {
	if subscriber.CreatedAt.IsZero() {
		subscriber.CreatedAt = time.Now()
	}
	subscriber.UpdatedAt = time.Now()

	type my Subscriber
	return bson.Marshal((*my)(subscriber))
}
