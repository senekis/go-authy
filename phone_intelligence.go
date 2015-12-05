package authy

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

type PhoneInfo struct {
  HttpResponse *http.Response
  Message      string `json:"message"`
  Type         string `json:"type"`
  Provider     string `json:"provider"`
  Ported       bool   `json:"ported"`
  Success      bool   `json:"success"`
}

func NewPhoneInfoRequest(response *http.Response) (*PhoneInfo, error) {
  phoneInfo := &PhoneInfo{HttpResponse: response}

  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)

  if err != nil {
    Logger.Println("Error reading from API:", err)
    return phoneInfo, err
  }

  err = json.Unmarshal(body, &phoneInfo)
  if err != nil {
    Logger.Println("Error parsing JSON:", err)
    return phoneInfo, err
  }

  return phoneInfo, nil
}
