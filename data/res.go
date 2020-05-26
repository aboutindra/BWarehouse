package data

type ResBool struct {
	Res bool `json:"res" bson:"res"`
}

type ResInt struct {
	Res int64 `json:"res" bson:"res"`
}

type ResArray struct {
	Res []interface{} `json:"res" bson:"res"`
}

type ResObj struct {
	Res interface{} `json:"res" bson:"res"`
}
