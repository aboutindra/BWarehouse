package main

import (
	"runtime"
	"ware/http2/master"

	"fmt"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {

	port := ":7554"

	runtime.GOMAXPROCS(4)

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/v1/api/all", master.GetAllDataMaster).Methods("GET")
	r.HandleFunc("/v1/api/total", master.GetTotalMaster).Methods("GET")
	r.HandleFunc("/v1/api/param/{id}", master.GetWithParamMaster).Methods("GET")

	r.HandleFunc("/v1/api/ino", master.InsertOneMaterial).Methods("POST")
	r.HandleFunc("/v1/api/inm", master.InsertManyMaterial).Methods("POST")

	r.HandleFunc("/v1/api/seto", master.PutOneMaster).Methods("PUT")

	r.HandleFunc("/v1/api/delo", master.DelOneMaster).Methods("POST")

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
