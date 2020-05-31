package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoDB) GetAll(col *mongo.Collection) (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	res, err := col.Find(ct, bson.M{})
	return res, err
}
func (m MongoDB) GetAllWithParam(col *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	res, err := col.Find(ct, filter)
	return res, err
}
func (m MongoDB) GetPagination(col *mongo.Collection, skip int64, lim int64) (*mongo.Cursor, error) {
	ct := m.MakeContext(10)
	option := options.FindOptions{}
	res, err := col.Find(ct, bson.M{}, option.SetLimit(lim), option.SetSkip(skip))
	return res, err

}
func (m MongoDB) GetOneWithParam(col *mongo.Collection, payload interface{}) *mongo.SingleResult {
	ct := m.MakeContext(10)
	tmp := col.FindOne(ct, payload)
	return tmp
}

func (m MongoDB) GetLen(col *mongo.Collection) (int64, error) {
	ct := m.MakeContext(10)
	tmp, err := col.CountDocuments(ct, bson.M{})
	return tmp, err
}

func (m MongoDB) GetLenWithParam(col *mongo.Collection, payload interface{}) (int64, error) {
	ct := m.MakeContext(10)
	tmp, err := col.CountDocuments(ct, payload)
	return tmp, err
}
