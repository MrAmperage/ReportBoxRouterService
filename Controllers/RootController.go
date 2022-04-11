package Controllers

import (
	"net/http"
	"text/template"
)

func RootController(ResponseWriter http.ResponseWriter, Request *http.Request) {

	Template, _ := template.ParseFiles("Static/Pages/Index.html")
	Template.Execute(ResponseWriter, nil)

}
