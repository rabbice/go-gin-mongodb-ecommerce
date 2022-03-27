package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	OrderedAt time.Time          `bson:"ordered_at" json:"ordered_at"`
}
