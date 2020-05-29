package output

import (
	"encoding/json"
	"net/http"
	"time"
	"ware/controller/ctro"
	"ware/http2"
)

func (o HttpOutput) InsertManyMaterial(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	var tmpReq []ctro.DataOutput
	json.NewDecoder(req.Body).Decode(&tmpReq)
	var tmpParam []interface{}
	len := len(tmpReq)
	var i = 0
	for i < len {
		var tmp ctro.DataOutput
		tmp = tmpReq[i]
		tmp.Tgl = time.Now()
		tmpParam = append(tmpParam, tmp)
		i++
	}
	err := d.InMany(tmpParam)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	json.NewEncoder(res).Encode(bol)
}
