package http2

import "net/http"

func SetHeader(res http.ResponseWriter) {

	res.Header().Add("content-type", "application/json")

}
