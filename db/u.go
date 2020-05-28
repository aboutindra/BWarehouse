package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongoDB) UpdOne(find interface{}, set primitive.D) error {
	ct := m.MakeContext(10)
	_, err := col.UpdateOne(ct, find, set)
	return err
}

func (m MongoDB) UpdMany(find interface{}, set primitive.D) error {
	ct := m.MakeContext(10)
	_, err := col.UpdateMany(ct, find, set)
	return err
}
