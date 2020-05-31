package output

import (
	"encoding/json"
	"net/http"
	"ware/controller/ctro"
	"ware/http2"
)

func (o HttpOutput) InsertManyMaterial(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmpReq ctro.DataOutArr
	json.NewDecoder(req.Body).Decode(&tmpReq)
	var tmpParam []interface{}
	len := len(tmpReq.Data)
	var i = 0
	for i < len {
		var tmp ctro.DataOutput
		tmp = tmpReq.Data[i]
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
