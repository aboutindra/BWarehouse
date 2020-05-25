package db

import "go.mongodb.org/mongo-driver/mongo"

func InOne(col mongo.Collection, payload interface{}) error {

	ct := MakeContext(10)

	_, err := col.InsertOne(ct, payload)

	return err

}

func InMany(col mongo.Collection, payload []interface{}) error {

	ct := MakeContext(25)

	_, err := col.InsertMany(ct, payload)

	return err

}
