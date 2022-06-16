package Controllers

import (
	"fmt"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
)

func RootController(ResponseWriter http.ResponseWriter, Request *http.Request, WebCore *WebCore.WebCore) (Data interface{}, Error error) {
	http.ServeFile(ResponseWriter, Request, fmt.Sprintf("%s/Pages/Index.html", WebCore.FileServer.StaticDirectory))

	return
}
