package ctro

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ControlOutput struct{}

func (c ControlOutput) ConvertCursor(cur *mongo.Cursor) []interface{} {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arr []interface{}
	var tmp DataOutput
	for cur.Next(ct) {
		cur.Decode(&tmp)
		arr = append(arr, tmp)
	}
	return arr
}

func (c ControlOutput) FormatToObj(cur *mongo.SingleResult) interface{} {
	var tmp DataOutput
	cur.Decode(&tmp)
	return tmp
}
