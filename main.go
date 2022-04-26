package main

import (
	"fmt"
	"os"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxRouterService/Controllers"
)

func main() {
	RouterService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := RouterService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
		os.Exit(0)
	}
	Controllers.InitControllers(RouterService)
	ErrorStartService := RouterService.StartRabbitMQ()
	if ErrorStartService != nil {

		fmt.Println(ErrorStartService)
		os.Exit(0)
	}

	ErrorStartWebServer := RouterService.StartWebServer()
	if ErrorStartWebServer != nil {
		fmt.Println(ErrorStartWebServer)
		os.Exit(0)
	}

}
