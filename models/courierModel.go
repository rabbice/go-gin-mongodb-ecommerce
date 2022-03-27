package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Courier struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `json:"name"`
	Code    uint16             `json:"code"`
	Partner bool               `json:"partner"`
}
