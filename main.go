package main

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func main() {
	RouterService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := RouterService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
	}
	ErrorStartService := RouterService.Start()
	if ErrorStartService != nil {

		fmt.Println(ErrorStartService)
	}

}
