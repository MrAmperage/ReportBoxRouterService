package Controllers

import (
	"net/http"
	"text/template"
)

func RootController(ResponseWriter http.ResponseWriter, Request *http.Request) (Data interface{}, Error error) {
	Template, Error := template.ParseFiles("Static/Pages/Index.html")
	if Error != nil {
		return
	}
	Error = Template.Execute(ResponseWriter, nil)
	if Error != nil {
		return
	}
	return
}
