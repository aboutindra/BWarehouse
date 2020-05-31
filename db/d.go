package db

import "go.mongodb.org/mongo-driver/mongo"

func (m MongoDB) DelOne(col *mongo.Collection, payload interface{}) error {
	ct := m.MakeContext(10)
	_, err := col.DeleteOne(ct, payload)
	return err
}
