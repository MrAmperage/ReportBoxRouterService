package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", RootController(ApplicationCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", AuthenticationController(ApplicationCore)).Methods("POST")
}
