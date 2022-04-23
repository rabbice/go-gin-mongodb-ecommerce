package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Delivery struct {
	ID      primitive.ObjectID `bson:"_id"`
	Weight  uint16             `bson:"package_weight" json:"package_weight"`
	Status  bool             `bson:"status" json:"status"`
	Courier []Courier          `bson:"courier" json:"courier"`
}
