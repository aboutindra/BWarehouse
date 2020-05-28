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

func UpdOneModel(col mongo.Collection, find interface{}, set data.ReqObjModel) error {

	update := bson.D{
		{"$set", bson.D{
			{"model", set.Model},
		}},
	}

	ct := MakeContext(10)

	_, err := col.UpdateOne(ct, find, update)

	return err

}

func UpdOneTipe(col mongo.Collection, find interface{}, set data.ReqObjTipe) error {

	update := bson.D{
		{"$set", bson.D{
			{"tipe", set.Tipe},
		}},
	}

	ct := MakeContext(10)

	_, err := col.UpdateOne(ct, find, update)

	return err

}

func UpdOneName(col mongo.Collection, find interface{}, set data.ReqObjName) error {

	update := bson.D{
		{"$set", bson.D{
			{"name", set.Name},
		}},
	}

	ct := MakeContext(10)

	_, err := col.UpdateOne(ct, find, update)

	return err

}

func UpdOneSub(col mongo.Collection, find interface{}, set data.ReqObjSub) error {

	update := bson.D{
		{"$set", bson.D{
			{"sub", set.Sub},
		}},
	}

	ct := MakeContext(10)

	_, err := col.UpdateOne(ct, find, update)

	return err

}
