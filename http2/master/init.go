package master

import (
	"ware/controller/m"
	"ware/db"
)

type HttpMaster struct{}

var d db.MongoDB
var c m.ControlMaster

func (h HttpMaster) init() {
	d = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Master"}
	c = m.ControlMaster{}
}
