package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Delivery struct {
	ID        primitive.ObjectID `bson:"_id"`
	Weight    uint16             `bson:"weight" json:"weight"`
	Status    bool               `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Courier   string             `bson:"courier" json:"courier"`
}
