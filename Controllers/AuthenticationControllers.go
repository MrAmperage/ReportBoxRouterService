package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
)

type AuthenticationData struct {
	Login    string
	Password string
}

func AuthenticationController(ApplicationCore *ApplicationCore.ApplicationCore) http.HandlerFunc {
	return func(ResponseWriter http.ResponseWriter, Request *http.Request) {

		var LoginPassword AuthenticationData
		ByteBody, Error := ioutil.ReadAll(Request.Body)
		if Error != nil {
			ResponseWriter.Write([]byte(Error.Error()))
		}
		Error = json.Unmarshal(ByteBody, &LoginPassword)
		if Error != nil {

			ResponseWriter.Write([]byte(Error.Error()))

		}
		ResponseWriter.Write([]byte(LoginPassword.Login))
	}

}
