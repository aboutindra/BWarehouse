package global

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"ware/controller/ctri"
	"ware/controller/ctrm"
	"ware/http2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (g HttpGlobal) GetMasterWithId(payload ctrm.ItemId) ctrm.DataMaster {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, _ := mongo.Connect(ct, options.Client().ApplyURI("mongodb://localhost:27017"))
	col := con.Database("WarehouseDB").Collection("Master")
	curr := col.FindOne(ct, payload)
	tmp := cm.FormatToObj(curr)
	con.Disconnect(ct)
	return tmp.(ctrm.DataMaster)
}

func (g HttpGlobal) AddStokMaster(id string, stok int64, wg *sync.WaitGroup) {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, _ := mongo.Connect(ct, options.Client().ApplyURI("mongodb://localhost:27017"))
	col := con.Database("WarehouseDB").Collection("Master")
	var find ctrm.ItemId
	find.Data = id
	format := cm.FormatUpdateStok(stok)
	col.UpdateOne(ct, find, format)
	con.Disconnect(ct)
	defer wg.Done()
}

func (g HttpGlobal) InInput(payload interface{}, wg *sync.WaitGroup) {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, _ := mongo.Connect(ct, options.Client().ApplyURI("mongodb://localhost:27017"))
	col := con.Database("WarehouseDB").Collection("Input")
	col.InsertOne(ct, payload)
	con.Disconnect(ct)
	defer wg.Done()
}

func (g HttpGlobal) MinusStokMaster(payload ctrm.ItemSub, len int, wg *sync.WaitGroup) {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, _ := mongo.Connect(ct, options.Client().ApplyURI("mongodb://localhost:27017"))
	col := con.Database("WarehouseDB").Collection("Master")
	var i int = 0
	for i < len {
		var find ctrm.ItemId
		find.Data = payload.Data[i].IdSub
		format := cm.FormatUpdateStok(payload.Data[i].Qty)
		col.UpdateOne(ct, find, format)
		i++
	}
	if i == len {
		con.Disconnect(ct)
		defer wg.Done()
	}
}

func (g HttpGlobal) InOut(payload []interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, _ := mongo.Connect(ct, options.Client().ApplyURI("mongodb://localhost:27017"))
	col := con.Database("WarehouseDB").Collection("Output")
	col.InsertMany(ct, payload)
}

func (g HttpGlobal) ToolsRecord(res http.ResponseWriter, req *http.Request) {

	http2.SetHeader(res)

	var filter ctrm.ItemId

	json.NewDecoder(req.Body).Decode(&filter)

	tmp := g.GetMasterWithId(filter)

	var tmpInput ctri.DataInput
	tmpInput.IdMate = tmp.IdMate
	tmpInput.Name = tmp.Name
	tmpInput.Qty = 1
	tmpInput.Tgl = time.Now()

	var tmpUpdateMany ctrm.ItemSub
	tmpUpdateMany.Data = tmp.Sub

	var leng int = len(tmp.Sub)

	var o int = 0

	var filter2 ctrm.ItemId

	var sub []interface{}

	for o < leng {
		sub = append(sub, tmp.Sub[o])
		sub[o] = tmp.Sub[o]
		filter2.Data = tmpUpdateMany.Data[o].IdSub
		tmp2 := g.GetMasterWithId(filter2)
		tmpUpdateMany.Data[o].Qty = tmp2.Stok - tmpUpdateMany.Data[o].Qty
		o++
	}

	if o == leng {

		var wg sync.WaitGroup

		wg.Add(4)

		go g.MinusStokMaster(tmpUpdateMany, leng, &wg)
		go g.AddStokMaster(tmp.IdMate, tmp.Stok+1, &wg)
		go g.InInput(tmpInput, &wg)
		go g.InOut(sub, &wg)

		wg.Wait()

		var tmpBool ctrm.ResBool

		tmpBool.Res = true
		json.NewEncoder(res).Encode(tmpBool)

	}

}
