package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email        *string            `json:"email" validate:"email,required"`
	Password     *string            `json:"password" validate:"required,min=8"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	Seller       *bool              `json:"seller" validate:"required"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
	Address      []Address          `bson:"address" json:"address"`
	Cart         []Cart             `bson:"usercart" json:"usercart"`
	OrderStatus  []Order            `bson:"orders" json:"orders"`
}
