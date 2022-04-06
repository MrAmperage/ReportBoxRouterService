module github.com/MrAmperage/ReportBoxRouterService

go 1.18

replace github.com/MrAmperage/GoWebStruct => ../GoWebStruct

require (
	github.com/MrAmperage/GoWebStruct v0.0.0-20220406024213-4ce96937be11
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
)
