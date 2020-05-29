package global

import (
	"ware/controller/ctri"
	"ware/controller/ctrm"
	"ware/controller/ctro"
	"ware/db"
)

type HttpGlobal struct{}

var dm db.MongoDB
var di db.MongoDB
var do db.MongoDB

var cm ctrm.ControlMaster
var ci ctri.ControlInput
var co ctro.ControlOutput

func init() {
	dm = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Master"}
	dm.Init()
	di = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Input"}
	di.Init()
	do = db.MongoDB{"mongodb://localhost:27017", "WarehouseDB", "Output"}
	do.Init()
	cm = ctrm.ControlMaster{}
	ci = ctri.ControlInput{}
	co = ctro.ControlOutput{}
}
