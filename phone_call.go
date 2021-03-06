package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PhoneCallRequest struct {
	HttpResponse *http.Response
	Message      string `json:"message"`
}

func NewPhoneCallRequest(response *http.Response) (*PhoneCallRequest, error) {
	smsRequest := &PhoneCallRequest{HttpResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return smsRequest, err
	}

	err = json.Unmarshal(body, &smsRequest)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return smsRequest, err
	}

	return smsRequest, nil
}

func (smsRequest *PhoneCallRequest) Valid() bool {
	if smsRequest.HttpResponse.StatusCode == 200 {
		return true
	}

	return false
}
