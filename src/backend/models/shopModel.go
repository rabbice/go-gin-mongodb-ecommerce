package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt   time.Time          `bson:"deleted_at" json:"deleted_at"`
	Shop_ID     string             `bson:"shop_id" json:"shop_id"`
}
