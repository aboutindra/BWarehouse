package input

import (
	"ware/controller/ctri"
	"ware/controller/ctrm"
	"ware/db"
)

type HttpInput struct{}

var c ctri.ControlInput
var d db.MongoDB

var bol ctrm.ResBool

func init() {
	d = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Input"}
	c = ctri.ControlInput{}
}
