package global

import (
	"encoding/json"
	"net/http"
	"time"
	"ware/controller/ctri"
	"ware/controller/ctrm"
	"ware/http2"
)

func (g HttpGlobal) GetMasterWithId(payload ctrm.ItemId) ctrm.DataMaster {
	curr := dm.GetOneWithParam(payload)
	tmp := cm.FormatToObj(curr)
	return tmp.(ctrm.DataMaster)
}

func (g HttpGlobal) AddStokMaster(id string, stok int64, sta1 chan bool) {
	var find ctrm.ItemId
	find.Data = id
	format := cm.FormatUpdateStok(stok)
	err := dm.UpdOne(find, format)
	if err != nil {
		sta1 <- false
	} else {
		sta1 <- true
	}
}

func (g HttpGlobal) InInput(payload interface{}, sta2 chan bool) {
	err := dm.InOne(payload)
	if err != nil {
		sta2 <- false
	} else {
		sta2 <- true
	}
}

func (g HttpGlobal) MinusStokMaster(payload ctrm.ItemSub, len int, sta3 chan bool) {
	var i int = 0
	for i < len {
		var find ctrm.ItemId
		find.Data = payload.Data[i].IdSub
		format := cm.FormatUpdateStok(payload.Data[i].Qty)
		dm.UpdOne(find, format)
		i++
	}
	sta3 <- true
}

func (g HttpGlobal) InOut(payload []interface{}, sta4 chan bool) {
	err := dm.InMany(payload)
	if err != nil {
		sta4 <- false
	} else {
		sta4 <- true
	}
}

func (g HttpGlobal) ToolsRecord(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var filter ctrm.ItemId

	json.NewDecoder(req.Body).Decode(&filter)

	tmp := g.GetMasterWithId(filter)

	var tmpInput ctri.DataInput
	tmpInput.IdMate = tmp.IdMate
	tmpInput.Name = tmp.Name
	tmpInput.Qty = 1
	tmpInput.Tgl = time.Now()

	var tmpUpdateMany ctrm.ItemSub
	tmpUpdateMany.Data = tmp.Sub

	var len int = len(tmp.Sub)

	var i int = 0

	var filter2 ctrm.ItemId

	for i < len {
		filter2.Data = tmpUpdateMany.Data[i].IdSub
		tmp2 := g.GetMasterWithId(filter2)
		tmpUpdateMany.Data[i].Qty = tmp2.Stok - tmpUpdateMany.Data[i].Qty
		i++
	}

	var tmpSub []interface{}

	tmpSub = append(tmpSub, tmp.Sub)

	sta1 := make(chan bool)
	sta2 := make(chan bool)
	sta3 := make(chan bool)
	sta4 := make(chan bool)

	go g.AddStokMaster(tmp.IdMate, tmp.Stok+1, sta1)
	go g.InInput(tmpInput, sta2)
	go g.MinusStokMaster(tmpUpdateMany, len, sta3)
	go g.InOut(tmpSub, sta4)

	x := <-sta1
	y := <-sta2
	z := <-sta3
	a := <-sta4

	var tmpBool ctrm.ResBool

	if x == true && y == true && z == true && a == true {
		tmpBool.Res = true
	} else {
		tmpBool.Res = false
	}

	json.NewEncoder(res).Encode(tmpBool)

}
