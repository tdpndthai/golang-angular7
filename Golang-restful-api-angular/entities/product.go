package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Price  float64       `json:"price" bson:"price"`
	Status bool          `json:"status" bson:"status"`
}
