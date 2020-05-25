package master

import (
	"encoding/json"
	"net/http"
	"ware/data"
	"ware/db"
	"ware/http2"
)

func PutOneMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection("mongodb://localhost:27017", "WarehouseDB", "Master")

	var tmpReq data.ReqObjUpdateStokDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.UpdOneStok(*col, tmpReq.F, tmpReq.S)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}
