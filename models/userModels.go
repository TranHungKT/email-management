package models

const (
	UserStatusEnabled  = "enabled"
	UserStatusDisabled = "disabled"
)

type UserBase struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	UserBase  `bson:",inline"`
	FirstName string `bson:"first_name" json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Status    string
}
