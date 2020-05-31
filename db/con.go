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

func (m MongoDB) MakeConnection() (*mongo.Collection, *mongo.Client) {

	ct := m.MakeContext(10)

	con, err := mongo.Connect(ct, options.Client().ApplyURI(m.Server))

	if err != nil {
		fmt.Println(err)
	}

	coll := con.Database(m.Dbname).Collection(m.Coll)

	return coll, con

}

func (m MongoDB) Dis(con *mongo.Client) {
	ct := m.MakeContext(10)
	con.Disconnect(ct)
}
