package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/ReportBoxRouterService/Models"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func GetManufacturers(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var Manufacturers []Models.Manufacturer
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Manufacturers", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		ReplyTo:       `amq.rabbitmq.reply-to`,
		CorrelationId: NewCorrelationId,
	})
	if Error != nil {
		return
	}
	RabbitMessage, Error := ReplySubscribe.GetMessageByCorrelationId(NewCorrelationId)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(RabbitMessage.Body, &Manufacturers)
	if Error != nil {
		return
	}
	return Manufacturers, Error
}
