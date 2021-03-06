package Controllers

import "github.com/MrAmperage/GoWebStruct/ApplicationCore"

func InitControllers(ApplicationCore *ApplicationCore.ApplicationCore) {

	ApplicationCore.WebCore.Router.HandleFunc("/", ApplicationCore.WebCore.Middleware.HandlerMiddleware(RootController, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/authentication", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AuthenticationController, &ApplicationCore.WebCore)).Methods("POST")
	//Пользователи
	ApplicationCore.WebCore.Router.HandleFunc("/api/Users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUser, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Users/{UserId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUser, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Users/{UserId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditUser, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Users", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUsers, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Users/{UserId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUser, &ApplicationCore.WebCore)).Methods("GET")
	//Конфигурация
	ApplicationCore.WebCore.Router.HandleFunc("/api/configuration/{Action}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetConfiguration, &ApplicationCore.WebCore)).Methods("GET")

	//Типы агрегатов
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes/{UnitTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUnitType, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUnitTypes, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUnitType, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes/{UnitTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUnitType, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitTypes/{UnitTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditUnitType, &ApplicationCore.WebCore)).Methods("PATCH")
	//Состояние агрегатов
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitStates", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUnitStates, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitStates/{UnitStateId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetUnitState, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitStates/{UnitStateId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteUnitState, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitStates/{UnitStateId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditUnitState, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/UnitStates", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddUnitState, &ApplicationCore.WebCore)).Methods("POST")
	//Типы транспорта
	ApplicationCore.WebCore.Router.HandleFunc("/api/TransportTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetTransportTypes, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/TransportTypes/{TransportTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetTransportType, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/TransportTypes/{TransportTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditTransportType, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/TransportTypes/{TransportTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteTransportType, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/TransportTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddTransportType, &ApplicationCore.WebCore)).Methods("POST")
	//Производители
	ApplicationCore.WebCore.Router.HandleFunc("/api/Manufacturers", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetManufacturers, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Manufacturers", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddManufacturer, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Manufacturers/{ManufacturerId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetManufacturer, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Manufacturers/{ManufacturerId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteManufacturer, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Manufacturers/{ManufacturerId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditManufacturer, &ApplicationCore.WebCore)).Methods("PATCH")
	//Организации
	ApplicationCore.WebCore.Router.HandleFunc("/api/Organizations", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetOrganizations, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Organizations/{OrganizationId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetOrganization, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Organizations/{OrganizationId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteOrganization, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Organizations/{OrganizationId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditOrganization, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Organizations", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddOrganization, &ApplicationCore.WebCore)).Methods("POST")
	//Типы грузов
	ApplicationCore.WebCore.Router.HandleFunc("/api/CargoTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetCargoTypes, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/CargoTypes/{CargoTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetCargoType, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/CargoTypes/{CargoTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteCargoType, &ApplicationCore.WebCore)).Methods("DELETE")
	ApplicationCore.WebCore.Router.HandleFunc("/api/CargoTypes/{CargoTypeId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditCargoType, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/CargoTypes", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddCargoType, &ApplicationCore.WebCore)).Methods("POST")
	//Роли
	ApplicationCore.WebCore.Router.HandleFunc("/api/Roles", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetRoles, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Roles/{RoleId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(GetRole, &ApplicationCore.WebCore)).Methods("GET")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Roles/{RoleId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(EditRole, &ApplicationCore.WebCore)).Methods("PATCH")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Roles", ApplicationCore.WebCore.Middleware.HandlerMiddleware(AddRole, &ApplicationCore.WebCore)).Methods("POST")
	ApplicationCore.WebCore.Router.HandleFunc("/api/Roles/{RoleId}", ApplicationCore.WebCore.Middleware.HandlerMiddleware(DeleteRole, &ApplicationCore.WebCore)).Methods("DELETE")
}
