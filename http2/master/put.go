package master

import (
	"encoding/json"
	"net/http"
	"ware/data"
	"ware/db"
	"ware/http2"
)

func PutOneStokMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

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

func PutOneTipeMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var tmpReq data.ReqObjUpdateTipeDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.UpdOneTipe(*col, tmpReq.F, tmpReq.S)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}

func PutOneModelMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var tmpReq data.ReqObjUpdateModelDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.UpdOneModel(*col, tmpReq.F, tmpReq.S)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}

func PutOneNameMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var tmpReq data.ReqObjUpdateNameDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.UpdOneName(*col, tmpReq.F, tmpReq.S)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}

func PutOneSubMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var tmpReq data.ReqObjUpdateSubDataMaster

	json.NewDecoder(req.Body).Decode(&tmpReq)

	errr := db.UpdOneSub(*col, tmpReq.F, tmpReq.S)

	var resBool data.ResBool

	if err != nil || errr != nil {
		resBool.Res = false
	} else {
		resBool.Res = true
	}

	json.NewEncoder(res).Encode(resBool)

}
