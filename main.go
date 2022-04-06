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
	RouterService.WebCore.Router.HandleFunc("/", Controllers.RootController)
	ErrorStartService := RouterService.Start()
	if ErrorStartService != nil {

		fmt.Println(ErrorStartService)
	}

}
