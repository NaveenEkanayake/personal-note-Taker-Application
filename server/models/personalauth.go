package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Fullname string             `json:"fullname" validate:"required,min=5,max=100"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required,min=5,max=100"`
}
