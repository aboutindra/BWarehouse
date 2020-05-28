package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqInt struct {
	Res int64
}

type SubMaterial struct {
	IdSub string `json:"idSub, omitempty" bson:"idSub, omitempty"`
	Name  string `json:"name, omitempty bson:"name, omitempty""`
	Qty   int64  `json:"qty, omitempty" bson:"qty, omitempty"`
}

type ReqDataMaster struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMaster string             `json:"idMate, omitempty" bson:"idMate, omitempty`
	Name     string             `json:"name, omitempty" bson:"name, omitempty"`
	Model    string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe     string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
	Stok     int64              `json:"stok, omitempty" bson:"stok, omitempty"`
	Sopname  int64              `json:"sopname, omitempty" bson:"sopname, omitempty"`
	Tgl      time.Time          `json:"tgl,omitempty" bson:"tgl,omitempty"`
	Sub      []SubMaterial      `json:"sub" bson:"sub"`
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

type ReqObjModel struct {
	Model string `json:"model, omitempty" bson:"model, omitempty"`
}

type ReqObjTipe struct {
	Tipe string `json:"tipe, omitempty" bson:"tipe, omitempty"`
}

type ReqObjSub struct {
	Sub []SubMaterial `json:"sub" bson:"sub"`
}

type ReqObjName struct {
	Name string `json:"name, omitempty" bson:"name, omitempty"`
}
type ReqObjUpdateModelDataMaster struct {
	F ReqItemIdDataMaster `json:"f" bson:"f"`
	S ReqObjModel         `json:"s" bson:"s"`
}
type ReqObjUpdateTipeDataMaster struct {
	F ReqItemIdDataMaster `json:"f" bson:"f"`
	S ReqObjTipe          `json:"s" bson:"s"`
}
type ReqObjUpdateSubDataMaster struct {
	F ReqItemIdDataMaster `json:"f" bson:"f"`
	S ReqObjSub           `json:"s" bson:"s"`
}
type ReqObjUpdateNameDataMaster struct {
	F ReqItemIdDataMaster `json:"f" bson:"f"`
	S ReqObjName          `json:"s" bson:"s"`
}
