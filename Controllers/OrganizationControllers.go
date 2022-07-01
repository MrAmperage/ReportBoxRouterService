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

func GetOrganizations(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var Organizations []Models.Organization
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Organizations", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &Organizations)
	if Error != nil {
		return
	}

	return []DataContainers.NaimedTable[Models.Organization]{{Table: Organizations, Name: "OrgatizationsTable", Type: "Table"}}, Error
}

func GetOrganization(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var Organization Models.Organization
	OrganizationId := mux.Vars(Request)["OrganizationId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Organizations", false, false, amqp.Publishing{
		Type:          "GET",
		ContentType:   "application/json",
		ReplyTo:       `amq.rabbitmq.reply-to`,
		CorrelationId: NewCorrelationId,
		Body:          []byte(OrganizationId),
	})
	if Error != nil {
		return
	}
	RabbitMessage, Error := ReplySubscribe.GetMessageByCorrelationId(NewCorrelationId)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(RabbitMessage.Body, &Organization)
	if Error != nil {
		return
	}

	return Organization, Error
}

func DeleteOrganization(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	OrganizationId := mux.Vars(Request)["OrganizationId"]
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Organizations", false, false, amqp.Publishing{
		Type:          "DELETE",
		Body:          []byte(OrganizationId),
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

func EditOrganization(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewOrganization Models.Organization
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {

		return
	}
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}

	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Organizations", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewOrganization)
	if Error != nil {
		return
	}

	return NewOrganization, Error
}

func AddOrganization(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreObject *WebCore.WebCore) (Data interface{}, Error error) {
	var NewOrganization Models.Organization
	NewCorrelationId := uuid.NewString()
	ReplySubscribe, Error := WebCoreObject.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("amq.rabbitmq.reply-to")
	if Error != nil {
		return
	}
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = WebCoreObject.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Organizations", false, false, amqp.Publishing{
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
	Error = json.Unmarshal(RabbitMessage.Body, &NewOrganization)
	if Error != nil {
		return
	}

	return NewOrganization, Error
}
