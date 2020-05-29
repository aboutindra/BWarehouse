package input

import (
	"encoding/json"
	"net/http"
	"ware/controller/ctri"
	"ware/http2"

	"github.com/gorilla/mux"
)

func (i HttpInput) GetAllData(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	cur, err := d.GetAll()
	tmp := c.ConvertCursor(cur)
	if err != nil || tmp == nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	json.NewEncoder(res).Encode(tmp)
}

func (i HttpInput) GetAllWithParamName(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	tmp := mux.Vars(req)["name"]
	obj := ctri.ItemName{tmp}
	cur, er := d.GetAllWithParam(obj)
	arr := c.ConvertCursor(cur)
	if arr == nil || er != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	json.NewEncoder(res).Encode(arr)
}

func (i HttpInput) GetAllWithParamId(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	tmp := mux.Vars(req)["id"]
	obj := ctri.ItemId{tmp}
	cur := d.GetOneWithParam(obj)
	has := c.FormatToObj(cur)
	if has == nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	json.NewEncoder(res).Encode(has)
}

func (i HttpInput) GetTotal(res http.ResponseWriter, req *http.Request) {
	http2.SetHeader(res)
	tot, err := d.GetLen()
	if err != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	obj := struct {
		Res int64 `json:"res"`
	}{tot}
	json.NewEncoder(res).Encode(obj)
}
