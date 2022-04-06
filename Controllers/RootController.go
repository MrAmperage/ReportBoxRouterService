package Controllers

import (
	"net/http"
)

func RootController(ResponseWriter http.ResponseWriter, Request *http.Request) {

	ResponseWriter.Write([]byte("123"))

}
