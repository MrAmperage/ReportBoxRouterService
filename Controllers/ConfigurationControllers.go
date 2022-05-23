package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func GetApplicationMenu(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	ApplicationMenu := new(map[string]interface{})
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Configuations", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		Body:          []byte("GetApplicationMenu"),
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
	Error = json.Unmarshal(RabbitMessage.Body, ApplicationMenu)
	if Error != nil {
		return
	}
	return ApplicationMenu, Error
}
