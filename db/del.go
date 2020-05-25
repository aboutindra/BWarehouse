package db

import "go.mongodb.org/mongo-driver/mongo"

func DelOne(col mongo.Collection, payload interface{}) error {

	ct := MakeContext(10)

	_, err := col.DeleteOne(ct, payload)

	return err

}
