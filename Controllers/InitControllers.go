package Controllers

import "github.com/MrAmperage/GoWebStruct/ApplicationCore"

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", ApplicationCore.WebCore.Middleware.HandlerMiddleware(RootController, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AuthenticationController, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUser, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users/{username}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUser, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users/{username}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditUser, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUsers, &ApplicationCore.WebCore)).Methods("GET")
}
