package main

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxRouterService/Controllers"
)

func main() {
	RouterService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := RouterService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
	}
	Controllers.InitControllers(RouterService)
	ErrorStartService := RouterService.StartRabbitMQ()
	if ErrorStartService != nil {

		fmt.Println(ErrorStartService)
	}
	ErrorStartWebServer := RouterService.StartWebServer()
	if ErrorStartWebServer != nil {
		fmt.Println(ErrorStartWebServer)
	}

}
