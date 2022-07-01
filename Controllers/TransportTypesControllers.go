package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/ReportBoxRouterService/DataContainers"
	"github.com/MrAmperage/ReportBoxRouterService/Models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func GetTransportTypes(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var TransportTypes []Models.TransportType
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "TransportTypes", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &TransportTypes)
	if Error != nil {
		return
	}
	return []DataContainers.ResponseContainer[[]Models.TransportType]{{Id: "TransportTypesTable", Container: TransportTypes, Type: "Table"}}, Error
}

func GetTransportType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var TransportType Models.TransportType
	TransportTypeId := mux.Vars(Request)["TransportTypeId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "TransportTypes", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		ReplyTo:       `amq.rabbitmq.reply-to`,
		CorrelationId: NewCorrelationId,
		Body:          []byte(TransportTypeId),
	})
	if Error != nil {
		return
	}
	RabbitMessage, Error := ReplySubscribe.GetMessageByCorrelationId(NewCorrelationId)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(RabbitMessage.Body, &TransportType)
	if Error != nil {
		return
	}

	return TransportType, Error
}
func EditTransportType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewTransportType Models.TransportType
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {

		return
	}
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "TransportTypes", false, false, amqp.Publishing{
		Type:          "PATCH",
		Body:          ByteBody,
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewTransportType)
	if Error != nil {
		return
	}

	return NewTransportType, Error
}

func DeleteTransportType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	TransportTypeId := mux.Vars(Request)["TransportTypeId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "TransportTypes", false, false, amqp.Publishing{
		Type:          "DELETE",
		Body:          []byte(TransportTypeId),
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

	return string(RabbitMessage.Body), Error
}

func AddTransportType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewTransportType Models.TransportType
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "TransportTypes", false, false, amqp.Publishing{
		Type:          "POST",
		Body:          ByteBody,
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewTransportType)
	if Error != nil {
		return
	}

	return NewTransportType, Error
}
