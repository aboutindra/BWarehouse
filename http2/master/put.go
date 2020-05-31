package master

import (
	"encoding/json"
	"net/http"
	"ware/controller/ctrm"
	"ware/http2"
)

func (h HttpMaster) PutOneStokMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.FormatUpdateStok
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateStok(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h HttpMaster) PutOneNameMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.FormatUpdateName
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateName(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h HttpMaster) PutOneModelMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.FormatUpdateModel
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateModel(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h HttpMaster) PutOneTipeMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.FormatUpdateTipe
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateTipe(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h HttpMaster) PutOneSubMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.FormatUpdateSub
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateSub(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
