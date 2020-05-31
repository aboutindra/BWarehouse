package master

import (
	"encoding/json"
	"net/http"
	"ware/controller/ctrm"
	"ware/http2"
)

func (h HttpMaster) DelOneMaster(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctrm.ItemId
	json.NewDecoder(req.Body).Decode(&tmp)
	er := d.DelOne(col, tmp)
	if er != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
