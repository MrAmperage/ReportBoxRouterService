package main

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxRouterService/Controllers"
	"github.com/streadway/amqp"
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
	ReplaySubscribe, Error := RouterService.WebCore.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		fmt.Println(Error)
	}
	Subscribe, Error := RouterService.WebCore.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("RouterQueue")
	if Error != nil {
		fmt.Println(Error)
	}
	go Subscribe.MessageProcessing(func(RabbitMQMessage amqp.Delivery) {
		fmt.Println(string(RabbitMQMessage.Body))

	})
	go ReplaySubscribe.MessageProcessing(func(RabbitMQMessage amqp.Delivery) {
		fmt.Println(string(RabbitMQMessage.Body))
	})
	ErrorStartWebServer := RouterService.StartWebServer()
	if ErrorStartWebServer != nil {
		fmt.Println(ErrorStartWebServer)
	}

}
