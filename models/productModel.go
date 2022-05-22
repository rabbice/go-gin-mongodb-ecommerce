package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `json:"name" validate:"required"`
	Description *string            `bson:"description" json:"description"`
	Price       *float64           `bson:"price" json:"price" validate:"required"`
	Image       *string            `bson:"image" json:"image"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at"`
}
