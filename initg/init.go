package main

import (
	"fmt"
	"os"
	"runtime"
	"ware/http2/global"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var htg global.HttpGlobal

func init() {
	htg = global.HttpGlobal{}
}

func main() {

	runtime.GOMAXPROCS(4)

	port := ":5554"

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.HandleFunc("/v1/api/record", htg.ToolsRecord).Methods("POST")

	r.StrictSlash(true)

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
