package main

import (
	"ware/http2/master"

	"fmt"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var htm master.HttpMaster

func init() {
	htm = master.HttpMaster{}
}

func main() {

	port := ":7554"

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/v1/api/all", htm.GetAllData).Methods("GET")
	r.HandleFunc("/v1/api/total", htm.GetTotal).Methods("GET")
	r.HandleFunc("/v1/api/param/i/{id}", htm.GetAllWithParamId).Methods("GET")
	r.HandleFunc("/v1/api/param/t/{tipe}", htm.GetAllWithParamTipe).Methods("GET")
	r.HandleFunc("/v1/api/param/m/{model}", htm.GetAllWithParamModel).Methods("GET")

	r.HandleFunc("/v1/api/ino", htm.InsertOneMaterial).Methods("POST")
	r.HandleFunc("/v1/api/inm", htm.InsertManyMaterial).Methods("POST")

	r.HandleFunc("/v1/api/seto/stok", htm.PutOneStokMaster).Methods("PUT")
	r.HandleFunc("/v1/api/seto/name", htm.PutOneNameMaster).Methods("PUT")
	r.HandleFunc("/v1/api/seto/model", htm.PutOneModelMaster).Methods("PUT")
	r.HandleFunc("/v1/api/seto/tipe", htm.PutOneTipeMaster).Methods("PUT")
	r.HandleFunc("/v1/api/seto/sub", htm.PutOneSubMaster).Methods("PUT")

	r.HandleFunc("/v1/api/delo", htm.DelOneMaster).Methods("DELETE")

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
