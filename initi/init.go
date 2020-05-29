package main

import (
	"fmt"
	"os"
	"ware/http2/input"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var hti input.HttpInput

func init() {
	hti = input.HttpInput{}
}

func main() {

	port := ":8554"

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/v1/api/all", hti.GetAllData).Methods("GET")
	r.HandleFunc("/v1/api/param/i/{id}", hti.GetAllWithParamId).Methods("GET")
	r.HandleFunc("/v1/api/param/n/{name}", hti.GetAllWithParamName).Methods("GET")
	r.HandleFunc("/v1/api/total", hti.GetTotal).Methods("GET")

	r.HandleFunc("/v1/api/ino", hti.InsertOneMaterial).Methods("POST")

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
