package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoDB) MakeContext(payload int) context.Context {
	duration := time.Duration(payload) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return ctx
}

func (m MongoDB) MakeConnection(payload string, db string, col string) *mongo.Collection {

	ct := m.MakeContext(10)

	con, err := mongo.Connect(ct, options.Client().ApplyURI(payload))

	if err != nil {
		fmt.Println(err)
	}

	coll := con.Database(db).Collection(col)

	return coll

}
