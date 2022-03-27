package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID          primitive.ObjectID `bson:"_id"`
	HouseNumber string             `json:"house_number"`
	Street      string             `json:"street"`
	City        string             `json:"city"`
	Zipcode     uint16             `json:"zipcode"`
	Country     string             `json:"country"`
	//User_ID     *string            `json:"user_id" validate:"required"`
}
