package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func GetConfiguration(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	ApplicationMenu := make(map[string]interface{})
	Action := mux.Vars(Request)["Action"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Configurations", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		Body:          []byte(Action),
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
	return ApplicationMenu, json.Unmarshal(RabbitMessage.Body, &ApplicationMenu)

}
