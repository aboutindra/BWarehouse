package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqInt struct {
	Res int64
}

type ReqDataMaster struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMaster string             `json:"idMate" bson:"idMate`
	Name     string             `json:"name" bson:"name"`
	Model    string             `json:"model" bson:"model"`
	Tipe     string             `json:"tipe" bson:"tipe"`
	Stok     int64              `json:"stok" bson:"stok"`
	Sopname  int64              `json:"sopname" bson:"sopname"`
	Tgl      time.Time          `json:"tgl" bson:"tgl"`
}

type ReqItemStokDataMaster struct {
	Stok int64 `json:"stok" bson:"stok"`
}

type ReqItemIdDataMaster struct {
	IdMaster string `json:"idMate" bson:"idMate`
}

type ReqObjUpdateStokDataMaster struct {
	F ReqItemIdDataMaster   `json:"f" bson:"f"`
	S ReqItemStokDataMaster `json:"s" bson:"s"`
}
