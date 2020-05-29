package main

import (
	"fmt"
	"os"
	"ware/http2/output"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var hto output.HttpOutput

func init() {
	hto = output.HttpOutput{}
}

func main() {

	port := ":9554"

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.HandleFunc("/v1/api/all", hto.GetAllData).Methods("GET")
	r.HandleFunc("/v1/api/param/i/{id}", hto.GetAllWithParamId).Methods("GET")
	r.HandleFunc("/v1/api/param/n/{name}", hto.GetAllWithParamName).Methods("GET")
	r.HandleFunc("/v1/api/total", hto.GetTotal).Methods("GET")

	r.HandleFunc("/v1/api/inm", hto.InsertManyMaterial).Methods("POST")

	r.StrictSlash(true)

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
