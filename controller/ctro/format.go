package ctro

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataOutput struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMate string             `json:"idMate, omitempty" bson:"idMate, omitempty"`
	Name   string             `json:"name, omitempty" bson:"name, omitempty"`
	Qty    int64              `json:"qty, omitempty" bson:"qty, omitempty"`
}

type DataOutArr struct {
	Data []DataOutput `json:"data" bson:"data"`
}

type ItemId struct {
	Data string `json:"idMate, omitempty" bson:"idMate, omitempty"`
}

type ItemName struct {
	Data string `json:"name, omitempty" bson:"name, omitempty"`
}

type ItemQty struct {
	Data int64 `json:"qty, omitempty" bson:"qty, omitempty"`
}

type ItemTime struct {
	Data time.Time `json:"tgl" bson:"tgl"`
}
