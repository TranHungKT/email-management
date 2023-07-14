package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ListTypePrivate = "private"
	ListTypePublic  = "public"

	ListOptinSingle = "single"
	ListOptinDouble = "double"
)

type List struct {
	Base `bson:",inline"`

	Name                  string `json:"name" bson:"name" validate:"required,max=20,min=3"`
	Type                  string
	Optin                 string
	Tags                  []string
	Description           string
	SubscriberCount       int                `json:"subscriber_count" bson:"subscriber_count"`
	SubscriberStatuses    map[string]int     `json:"subscriber_statuses" bson:"subscriber_statuses"`
	SubscriberID          int                `json:"subscriber_id" bson:"subscriber_id"`
	SubscriptionStatus    string             `json:"subscription_status" bson:"subscription_status,omitempty"`
	SubscriptionCreatedAt primitive.DateTime `json:"subscription_created_at" bson:"subscription_created_at,omitempty"`
	SubscriptionUpdatedAt primitive.DateTime `json:"subscription_updated_at" bson:"subscription_updated_at,omitempty"`

	Total int
}

func (list *List) MarshalBSON() ([]byte, error) {
	if list.CreatedAt.IsZero() {
		list.CreatedAt = time.Now()
	}
	list.UpdatedAt = time.Now()

	type my List
	return bson.Marshal((*my)(list))
}
