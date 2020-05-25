package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeContext(payload int) context.Context {
	duration := time.Duration(payload) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return ctx
}

func MakeConnection(payload string, db string, col string) (*mongo.Collection, error) {

	ct := MakeContext(10)

	con, err := mongo.Connect(ct, options.Client().ApplyURI(payload))

	if err != nil {
		return nil, err
	}

	coll := con.Database(db).Collection(col)

	return coll, nil

}
