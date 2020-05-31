package master

import (
	"encoding/json"
	"net/http"
	"time"
	"ware/controller/ctrm"
	"ware/http2"
)

func (h HttpMaster) InsertOneMaterial(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.DataMaster
	json.NewDecoder(req.Body).Decode(&tmp)
	tmp.Tgl = time.Now()
	er := d.InOne(col, tmp)
	if er != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h HttpMaster) InsertManyMaterial(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmpReq []ctrm.DataMaster
	json.NewDecoder(req.Body).Decode(&tmpReq)
	var tmpParam []interface{}
	len := len(tmpReq)
	i := 0
	for i < len {
		var tmp ctrm.DataMaster
		tmp = tmpReq[i]
		tmp.Tgl = time.Now()
		tmpParam = append(tmpParam, tmp)
		i++
	}
	err := d.InMany(col, tmpParam)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
