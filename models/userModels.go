package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	UserTypeSuperAdmin = "superAdmin"
	UserTypeUser       = "user"
	UserStatusEnabled  = "enabled"
	UserStatusDisabled = "disabled"
)

type UserBase struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Base      `bson:",inline"`
	UserBase  `bson:",inline"`
	FirstName string `bson:"first_name" json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Type      string
	Status    string
}

func (user *User) MarshalBSON() ([]byte, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(user))
}
