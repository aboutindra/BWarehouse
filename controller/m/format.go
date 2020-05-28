package m

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataSub struct {
	IdSub string `json:"idSub, omitempty" bson:"idSub, omitempty"`
	Name  string `json:"name, omitempty bson:"name, omitempty""`
	Qty   int64  `json:"qty, omitempty" bson:"qty, omitempty"`
}

type DataMaster struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMaster string             `json:"idMate, omitempty" bson:"idMate, omitempty`
	Name     string             `json:"name, omitempty" bson:"name, omitempty"`
	Model    string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe     string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
	Stok     int64              `json:"stok, omitempty" bson:"stok, omitempty"`
	Sopname  int64              `json:"sopname, omitempty" bson:"sopname, omitempty"`
	Tgl      time.Time          `json:"tgl,omitempty" bson:"tgl,omitempty"`
	Sub      []DataSub          `json:"sub" bson:"sub"`
}

type ItemId struct {
	Data string `json:"idMate, omitempty" bson:"idMate, omitempty`
}

type ItemName struct {
	Data string `json:"name, omitempty" bson:"name, omitempty"`
}

type ItemModel struct {
	Data string `json:"model, omitempty" bson:"model, omitempty"`
}

type ItemTipe struct {
	Data string `json:"tipe, omitempty" bson:"tipe,omitempty"`
}

type ItemStok struct {
	Data int64 `json:"stok" bson:"stok"`
}

type ItemSub struct {
	Data []SubMaterial `json:"sub" bson:"sub"`
}

type FormatUpdateName struct {
	Find ItemId   `json:"find" bson:"find"`
	Set  ItemName `json:"set" bson:"set"`
}

type FormatUpdateModel struct {
	Find ItemId    `json:"find" bson:"find"`
	Set  ItemModel `json:"set" bson:"set"`
}

type FormatUpdateTipe struct {
	Find ItemId   `json:"find" bson:"find"`
	Set  ItemTipe `json:"set" bson:"set"`
}

type FormatUpdateStok struct {
	Find ItemId   `json:"find" bson:"find"`
	Set  ItemStok `json:"set" bson:"set"`
}

type FormatUpdateSub struct {
	Find ItemId  `json:"find" bson:"find"`
	Set  ItemSub `json:"set" bson:"set"`
}
