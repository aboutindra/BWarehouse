package db

import (
	"ware/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type F struct {
	Col  mongo.Collection
	Tipe int
}

var ret data.ResDataMaster
var ret2 data.ResDataIO

func (m F) GetAllData() ([]interface{}, error) {

	var tmpRes []interface{}

	ct := MakeContext(60)
	ct2 := MakeContext(10)

	res, err := m.Col.Find(ct, bson.M{})

	if err != nil {
		return nil, err
	}

	for res.Next(ct2) {

		if m.Tipe == 1 {
			res.Decode(&ret)
			tmpRes = append(tmpRes, ret)
		} else {
			res.Decode(&ret2)
			tmpRes = append(tmpRes, ret2)
		}

	}

	return tmpRes, nil

}

func (m F) GetAllWithParam(filter interface{}) ([]interface{}, error) {

	var tmpRes []interface{}

	ct := MakeContext(10)

	res, err := m.Col.Find(ct, filter)

	if err != nil {
		return nil, err
	}

	for res.Next(ct) {

		if m.Tipe == 1 {
			res.Decode(&ret)
			tmpRes = append(tmpRes, ret)
		} else {
			res.Decode(&ret2)
			tmpRes = append(tmpRes, ret2)
		}

	}

	return tmpRes, nil
}

func (m F) GetPagination(lim int64, skip int64) ([]interface{}, error) {

	var tmpRes []interface{}

	ct := MakeContext(10)

	option := options.FindOptions{}

	res, err := m.Col.Find(ct, bson.M{}, option.SetLimit(lim), option.SetSkip(skip))

	if err != nil {
		return nil, err
	}

	for res.Next(ct) {

		if m.Tipe == 1 {
			res.Decode(&ret)
			tmpRes = append(tmpRes, ret)
		} else {
			res.Decode(&ret2)
			tmpRes = append(tmpRes, ret2)
		}

	}

	return tmpRes, nil

}

func (m F) GetWithParam(payload interface{}) interface{} {

	ct := MakeContext(10)

	if m.Tipe == 1 {

		m.Col.FindOne(ct, payload).Decode(&ret)

		if ret.Id.IsZero() == false {
			return ret
		} else {
			return nil
		}

	} else {

		m.Col.FindOne(ct, payload).Decode(&ret2)

		if ret2.Id.IsZero() == false {
			return ret
		} else {
			return nil
		}

	}

}

// ------ Universal ----

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
