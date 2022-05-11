package Controllers

import "github.com/MrAmperage/GoWebStruct/ApplicationCore"

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", ApplicationCore.WebCore.Middleware.HandlerMiddleware(RootController, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AuthenticationController, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(UsersController, &ApplicationCore.WebCore)).Methods("POST")
}
