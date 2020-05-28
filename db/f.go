package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoDB) GetAll() (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	res, err := col.Find(ct, bson.M{})
	return res, err
}
func (m MongoDB) GetAllWithParam(filter interface{}) (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	res, err := col.Find(ct, filter)
	return res, err
}
func (m MongoDB) GetPagination(skip int64, lim int64) (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	option := options.FindOptions{}
	res, err := col.Find(ct, bson.M{}, option.SetLimit(lim), option.SetSkip(skip))
	return res, err

}
func (m MongoDB) GetOneWithParam(payload interface{}) interface{} {
	var tmp interface{}
	ct := m.MakeContext(10)
	col.FindOne(ct, payload).Decode(&tmp)
	return tmp
}

func (m MongoDB) GetLen() (int64, error) {
	ct := m.MakeContext(10)
	tmp, err := col.CountDocuments(ct, bson.M{})
	return tmp, err
}

func (m MongoDB) GetLenWithParam(payload interface{}) (int64, error) {
	ct := m.MakeContext(10)
	tmp, err := col.CountDocuments(ct, payload)
	return tmp, err
}
