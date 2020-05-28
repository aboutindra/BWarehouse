package master

import (
	"encoding/json"
	"net/http"
	"time"
	"ware/data"
	"ware/db"
	"ware/http2"
)

func InsertOneMaterial(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var tmpReq data.ReqDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	tmpReq.Tgl = time.Now()

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var resBool data.ResBool

	errr := db.InOne(*col, tmpReq)

	if err != nil || errr != nil {
		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

	} else {
		resBool.Res = true
		json.NewEncoder(res).Encode(resBool)
	}

}

func InsertManyMaterial(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var tmpReq []interface{}

	json.NewDecoder(req.Body).Decode(&tmpReq)

	var tmpParam []interface{}

	var len int = len(tmpReq)

	var i int = 0

	for i < len {

		var tmp data.ReqDataMaster

		tmp = tmpReq[i].(data.ReqDataMaster)

		tmp.Tgl = time.Now()

		tmpParam = append(tmpParam, tmp)

		i++
	}

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var resBool data.ResBool

	errr := db.InMany(*col, tmpParam)

	if err != nil || errr != nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	} else {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return
	}

}
