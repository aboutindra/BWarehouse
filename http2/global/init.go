package global

import (
	"ware/controller/ctri"
	"ware/controller/ctrm"
	"ware/controller/ctro"
)

type HttpGlobal struct{}

var cm ctrm.ControlMaster
var ci ctri.ControlInput
var co ctro.ControlOutput

func init() {
	cm = ctrm.ControlMaster{}
	ci = ctri.ControlInput{}
	co = ctro.ControlOutput{}
}
