package Controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/streadway/amqp"
)

func AuthenticationController(ResponseWriter http.ResponseWriter, Request *http.Request, WebCore *WebCore.WebCore) (Data interface{}, Error error) {
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = WebCore.RabbitMQ.RabbitMQChanel.Chanel.Publish("RportBoxExchange", "Authentication", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        ByteBody,
	})
	if Error != nil {
		return
	}

	return "ok", nil

}
