package m

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ControlMaster struct{}

func (c ControlMaster) ConvertCursor(cur *mongo.Cursor) []interface{} {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arr []interface{}
	var tmp DataMaster
	for cur.Next(ct) {
		cur.Decode(&tmp)
		arr = append(arr, tmp)
	}
	return arr
}

func (c ControlMaster) FormatUpdateStok(payload int64) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"stok", payload},
		}},
	}
	return tmp
}

func (c ControlMaster) FormatUpdateName(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"name", payload},
		}},
	}
	return tmp
}

func (c ControlMaster) FormatUpdateModel(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"model", payload},
		}},
	}
	return tmp
}

func (c ControlMaster) FormatUpdateTipe(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"tipe", payload},
		}},
	}
	return tmp
}

func (c ControlMaster) FormatUpdateSub(payload []DataSub) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"sub", payload},
		}},
	}
	return tmp
}
