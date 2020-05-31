package db

import "go.mongodb.org/mongo-driver/mongo"

func (m MongoDB) InOne(col *mongo.Collection, payload interface{}) error {
	ct := m.MakeContext(10)
	_, err := col.InsertOne(ct, payload)
	return err
}

func (m MongoDB) InMany(col *mongo.Collection, payload []interface{}) error {
	ct := m.MakeContext(25)
	_, err := col.InsertMany(ct, payload)
	return err
}
