package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore"
)

type AuthenticationData struct {
	Login    string
	Password string
}

func AuthenticationController(ResponseWriter http.ResponseWriter, Request *http.Request, WebCore WebCore.WebCore) (Data interface{}, Error error) {
	var LoginPassword AuthenticationData
	ByteBody, Error := ioutil.ReadAll(Request.Body)
	if Error != nil {
		return
	}
	Error = json.Unmarshal(ByteBody, &LoginPassword)
	if Error != nil {

		return

	}
	Data = LoginPassword.Login
	return

}
