package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/ReportBoxRouterService/Models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func GetCargoTypes(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var CargoTypes []Models.CargoType
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "CargoTypes", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &CargoTypes)
	if Error != nil {
		return
	}
	return CargoTypes, Error
}

func GetCargoType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var CargoType Models.CargoType
	CargoTypeId := mux.Vars(Request)["CargoTypeId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "CargoTypes", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		ReplyTo:       `amq.rabbitmq.reply-to`,
		CorrelationId: NewCorrelationId,
		Body:          []byte(CargoTypeId),
	})
	if Error != nil {
		return
	}
	RabbitMessage, Error := ReplySubscribe.GetMessageByCorrelationId(NewCorrelationId)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(RabbitMessage.Body, &CargoType)
	if Error != nil {
		return
	}

	return CargoType, Error
}

func DeleteCargoType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	CargoTypeId := mux.Vars(Request)["CargoTypeId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "CargoTypes", false, false, amqp.Publishing{
		Type:          "DELETE",
		Body:          []byte(CargoTypeId),
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

func EditCargoType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewCargoType Models.CargoType
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {

		return
	}
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "CargoTypes", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewCargoType)
	if Error != nil {
		return
	}

	return NewCargoType, Error
}

func AddCargoType(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewCargoType Models.CargoType
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "CargoTypes", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewCargoType)
	if Error != nil {
		return
	}

	return NewCargoType, Error
}
