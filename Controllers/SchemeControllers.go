package Controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func AddScheme(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	NewCorrelationId := uuid.NewString()

	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Schemes", false, false, amqp.Publishing{
		Type:          "POST",
		ContentType:   "application/json",
		Body:          ByteBody,
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
	return string(RabbitMessage.Body), Error
}
