package output

import (
	"ware/controller/ctrm"
	"ware/controller/ctro"
	"ware/db"
)

type HttpOutput struct{}

var c ctro.ControlOutput
var d db.MongoDB

var bol ctrm.ResBool

func init() {
	c = ctro.ControlOutput{}
	d = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Output"}
	d.Init()
}
