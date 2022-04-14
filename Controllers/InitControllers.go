package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", ApplicationCore.WebCore.Middleware.ErrorHandlerMiddleware(RootController)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", ApplicationCore.WebCore.Middleware.ErrorHandlerMiddleware(AuthenticationController)).Methods("POST")
}
