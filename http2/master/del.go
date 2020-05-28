package master

import (
	"encoding/json"
	"net/http"
	"ware/data"
	"ware/db"
	"ware/http2"
)

func DelOneMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var tmpReq data.ReqItemIdDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.DelOne(*col, tmpReq)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}
