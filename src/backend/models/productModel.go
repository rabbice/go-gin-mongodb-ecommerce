package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `json:"name" validate:"required"`
	Description string             `json:"description"`
	Price       float64            `json:"price" validate:"required"`
	Image       string             `json:"image"`
	Tags        []string           `json:"tags"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   time.Time          `json:"deleted_at"`
	Product_ID  string             `json:"product_id"`
	Shop_ID     *string            `json:"shop_id" validate:"required"`
}
