package master

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ware/data"
	"ware/db"
	"ware/http2"

	"github.com/gorilla/mux"
)

const (
	mongoDB string = "mongodb://localhost:27017"
	dbName  string = "WarehouseDB"
	coll    string = "Master"
)

var m db.F

func init() {

	col, _ := db.MakeConnection(mongoDB, dbName, coll)
	m = db.F{*col, 1}

}

func GetAllDataMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var resBool data.ResBool

	result, err := m.GetAllData()

	var resArr data.ResArray

	if result == nil || err != nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	resArr.Res = result

	json.NewEncoder(res).Encode(resArr)

}

func GetAllWithParamModel(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var tmpReq data.ReqObjModel

	tmpReq.Model = mux.Vars(req)["model"]

	result, err := m.GetAllWithParam(tmpReq)

	var tmpBool data.ResBool

	if err != nil || result == nil {

		tmpBool.Res = false
		json.NewEncoder(res).Encode(tmpBool)

		return

	}

	var tmpArr data.ResArray

	tmpArr.Res = result

	json.NewEncoder(res).Encode(tmpArr)

}

func GetAllWithParamTipe(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var tmpReq data.ReqObjTipe

	tmpReq.Tipe = mux.Vars(req)["tipe"]

	result, err := m.GetAllWithParam(tmpReq)

	var tmpBool data.ResBool

	if err != nil || result == nil {

		tmpBool.Res = false
		json.NewEncoder(res).Encode(tmpBool)

		return

	}

	var tmpArr data.ResArray

	tmpArr.Res = result

	json.NewEncoder(res).Encode(tmpArr)

}

func GetWithParamMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var resBool data.ResBool

	var tmpFormat data.ReqItemIdDataMaster

	tmpReq := mux.Vars(req)["id"]

	tmpFormat.IdMaster = tmpReq

	result := m.GetWithParam(tmpFormat)

	var resObj data.ResObj

	if result == nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	resObj.Res = result

	json.NewEncoder(res).Encode(resObj)

}

func GetTotalMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection(mongoDB, dbName, coll)

	var resBool data.ResBool

	var tmpRes data.ResInt

	var tmpInt int64

	tmpInt, errr := db.GetLen(*col)

	tmpRes.Res = tmpInt

	if errr != nil || err != nil {

		fmt.Println(errr)

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	json.NewEncoder(res).Encode(tmpRes)

}
