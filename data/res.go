package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResBool struct {
	Res bool `json:"res" bson:"res"`
}

type ResInt struct {
	Res int64 `json:"res" bson:"res"`
}

type ResArray struct {
	Res []interface{} `json:"res" bson:"res"`
}

type ResObj struct {
	Res interface{} `json:"res" bson:"res"`
}

type ResSubMaterial struct {
	IdSub string `json:"idSub, omitempty" bson:"idSub, omitempty"`
	Name  string `json:"name, omitempty bson:"name, omitempty""`
	Qty   int64  `json:"qty, omitempty" bson:"qty, omitempty"`
}

type ResDataMaster struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMaster string             `json:"idMate, omitempty" bson:"idMate, omitempty`
	Name     string             `json:"name, omitempty" bson:"name, omitempty"`
	Model    string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe     string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
	Stok     int64              `json:"stok, omitempty" bson:"stok, omitempty"`
	Sopname  int64              `json:"sopname, omitempty" bson:"sopname, omitempty"`
	Tgl      time.Time          `json:"tgl,omitempty" bson:"tgl,omitempty"`
	Sub      []ResSubMaterial   `json:"sub" bson:"sub"`
}

type ResDataIO struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMaster string             `json:"idMate, omitempty" bson:"idMate, omitempty`
	Qty      int64              `json:"qty, omitempty" bson:"qty, omitempty"`
	Tgl      time.Time          `json:"tgl,omitempty" bson:"tgl,omitempty"`
}
