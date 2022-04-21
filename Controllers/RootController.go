package Controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/MrAmperage/GoWebStruct/WebCore"
)

func RootController(ResponseWriter http.ResponseWriter, Request *http.Request, WebCore *WebCore.WebCore) (Data interface{}, Error error) {
	Template, Error := template.ParseFiles(fmt.Sprintf("%s/Pages/Index.html", WebCore.FileServer.StaticDirectory))
	if Error != nil {
		return
	}
	Error = Template.Execute(ResponseWriter, nil)
	if Error != nil {
		return
	}
	return
}
