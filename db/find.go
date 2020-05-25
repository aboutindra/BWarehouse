package db

import (
	"ware/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllData(col mongo.Collection) ([]interface{}, error) {

	var tmpRes []interface{}

	ct := MakeContext(60)
	ct2 := MakeContext(10)

	res, err := col.Find(ct, bson.M{})

	if err != nil {
		return nil, err
	}

	for res.Next(ct2) {

		var tmp data.ReqDataMaster

		res.Decode(&tmp)

		tmpRes = append(tmpRes, tmp)

	}

	return tmpRes, nil

}

func GetPagination(col mongo.Collection, lim int64, skip int64) ([]interface{}, error) {

	var tmpRes []interface{}

	ct := MakeContext(10)

	option := options.FindOptions{}

	res, err := col.Find(ct, bson.M{}, option.SetLimit(lim), option.SetSkip(skip))

	if err != nil {
		return nil, err
	}

	for res.Next(ct) {

		var tmp data.ReqDataMaster

		res.Decode(&tmp)

		tmpRes = append(tmpRes, tmp)

	}

	return tmpRes, nil

}

func GetWithParam(col mongo.Collection, payload interface{}) interface{} {

	var tmpRes data.ReqDataMaster

	ct := MakeContext(10)

	col.FindOne(ct, payload).Decode(&tmpRes)

	return tmpRes

}

func GetLen(col mongo.Collection) (int64, error) {

	var tmpLen int64

	ct := MakeContext(10)

	tmpLen, err := col.CountDocuments(ct, bson.M{})

	if err != nil {
		return 0, err
	}

	return tmpLen, nil

}

func GetLenWithParam(col mongo.Collection, payload interface{}) (int64, error) {

	var tmpLen int64

	ct := MakeContext(10)

	tmpLen, err := col.CountDocuments(ct, payload)

	if err != nil {
		return 0, err
	}

	return tmpLen, nil

}
