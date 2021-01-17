package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	First   string             `json:”first”   bson:"first"   validate:"required"`
	Last    string             `json:”last”    bson:"last"    validate:"required"`
	Age     int                `json:”age”     bson:"age"     validate:"required"`
	City    string             `json:”city”    bson:"city"    validate: "required"`
	Country string             `json:”country” bson:"country" validate: "required"`
}
type Users []User

type UserValidator struct {
	Validator *validator.Validate
}

func (u *UserValidator) Validate(i interface{}) error {
	return u.Validator.Struct(i)
}
