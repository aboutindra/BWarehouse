package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Server string
	Dbname string
	Coll   string
}

var col *mongo.Collection

func (m MongoDB) Init() {
	col = m.MakeConnection(m.Server, m.Dbname, m.Coll)
}
