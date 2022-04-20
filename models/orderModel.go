package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `bson:"name" json:"name"`
	Description *string            `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Image       *string            `bson:"image" json:"image"`
}

type Order struct {
	ID            primitive.ObjectID `bson:"_id"`
	Cart          []Cart             `bson:"order_list" json:"order_list"`
	OrderedAt     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price         float64            `bson:"total_price" json:"total_price"`
	PaymentMethod Payment            `bson:"payment_method" json:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
