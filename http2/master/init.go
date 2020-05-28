package master

import (
	"ware/controller/ctrm"
	"ware/db"
)

type HttpMaster struct{}

var d db.MongoDB
var c ctrm.ControlMaster

func init() {
	d = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Master"}
	d.Init()
	c = ctrm.ControlMaster{}
}
