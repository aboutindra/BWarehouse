package db

import (
	"ware/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdOneStok(col mongo.Collection, find interface{}, set data.ReqItemStokDataMaster) error {

	update := bson.D{
		{"$set", bson.D{
			{"stok", set.Stok},
		}},
	}

	ct := MakeContext(10)

	_, err := col.UpdateOne(ct, find, update)

	return err

}
