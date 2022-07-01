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
	return []DataContainers.ResponseContainer[[]Models.Manufacturer]{{Id: "ManufacturersTable", Container: Manufacturers, Type: "Table"}}, Error
}

func EditManufacturer(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewManufacturer Models.Manufacturer
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {

		return
	}
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Manufacturers", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewManufacturer)
	if Error != nil {
		return
	}

	return NewManufacturer, Error
}

func GetManufacturer(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var Manufacturer Models.Manufacturer
	TransportTypeId := mux.Vars(Request)["ManufacturerId"]
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
		Body:          []byte(TransportTypeId),
	})
	if Error != nil {
		return
	}
	RabbitMessage, Error := ReplySubscribe.GetMessageByCorrelationId(NewCorrelationId)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(RabbitMessage.Body, &Manufacturer)
	if Error != nil {
		return
	}

	return Manufacturer, Error
}

func DeleteManufacturer(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	ManufacturerId := mux.Vars(Request)["ManufacturerId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Manufacturers", false, false, amqp.Publishing{
		Type:          "DELETE",
		Body:          []byte(ManufacturerId),
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

func AddManufacturer(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewManufacturer Models.Manufacturer
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Manufacturers", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewManufacturer)
	if Error != nil {
		return
	}

	return NewManufacturer, Error
}
