package main

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

func main() {
	RouterService := &ApplicationCore.ApplicationCore{}
	Error := RouterService.Start()
	if Error != nil {

		fmt.Println(Error)
	}
}
