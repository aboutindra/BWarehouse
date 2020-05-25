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

func GetAllDataMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection("mongodb://localhost:27017", "WarehouseDB", "Master")

	var resBool data.ResBool

	if err != nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	result, errr := db.GetAllData(*col)

	if errr != nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	json.NewEncoder(res).Encode(result)

}

func GetWithParamMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection("mongodb://localhost:27017", "WarehouseDB", "Master")

	var resBool data.ResBool

	if err != nil {

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	var tmpFormat data.ReqItemIdDataMaster

	tmpReq := mux.Vars(req)["id"]

	tmpFormat.IdMaster = tmpReq

	result := db.GetWithParam(*col, tmpFormat)

	json.NewEncoder(res).Encode(result)

}

func GetTotalMaster(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	col, err := db.MakeConnection("mongodb://localhost:27017", "WarehouseDB", "Master")

	var resBool data.ResBool

	if err != nil {

		fmt.Println(err)

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	var tmpRes data.ResInt

	var tmpInt int64

	tmpInt, errr := db.GetLen(*col)

	tmpRes.Res = tmpInt

	if errr != nil {

		fmt.Println(errr)

		resBool.Res = false
		json.NewEncoder(res).Encode(resBool)

		return

	}

	json.NewEncoder(res).Encode(tmpRes)

}
