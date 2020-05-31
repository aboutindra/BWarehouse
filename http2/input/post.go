package input

import (
	"encoding/json"
	"net/http"
	"time"
	"ware/controller/ctri"
	"ware/http2"
)

func (i HttpInput) InsertOneMaterial(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	col, con := d.MakeConnection()
	var tmp ctri.DataInput
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
