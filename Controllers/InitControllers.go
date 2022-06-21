package Controllers

import "github.com/MrAmperage/GoWebStruct/ApplicationCore"

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", ApplicationCore.WebCore.Middleware.HandlerMiddleware(RootController, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AuthenticationController, &ApplicationCore.WebCore)).Methods("POST")
	//Пользователи
	ApplicationCore.WebCore.Router.HandleFunc("/api/users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUser, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users/{Username}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUser, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users/{Username}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditUser, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUsers, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/users/{Username}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUser, &ApplicationCore.WebCore)).Methods("GET")
	//Схемы
	ApplicationCore.WebCore.Router.HandleFunc("/api/schemes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddScheme, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/schemes/{SchemeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteScheme, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/schemes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetSchemes, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/schemes/{SchemeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditScheme, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/schemes/{SchemeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetScheme, &ApplicationCore.WebCore)).Methods("GET")

	//Конфигурация
	ApplicationCore.WebCore.Router.HandleFunc("/api/configuration/{Action}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetConfiguration, &ApplicationCore.WebCore)).Methods("GET")

	//Типы агрегатов
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUnitTypes, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUnitType, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes/{UnitTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUnitType, &ApplicationCore.WebCore)).Methods("DELETE")

}
