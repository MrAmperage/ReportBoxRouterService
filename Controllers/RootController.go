package Controllers

import (
	"net/http"
	"text/template"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func RootController(ApplicationCore *ApplicationCore.ApplicationCore) http.HandlerFunc {
	return func(ResponseWriter http.ResponseWriter, Request *http.Request) {
		Template, _ := template.ParseFiles("Static/Pages/Index.html")
		Template.Execute(ResponseWriter, nil)
	}

}
