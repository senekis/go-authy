package authy

import (
  "net/url"
  "testing"
)

func Test_RequestPhoneInformation(t *testing.T) {
  api := NewSandboxAuthyApi("bf12974d70818a08199d17d5e2bae630")

  phoneResponse, err := api.PhoneIntelligence(1, "775-461-5609", url.Values{})

  if err == nil {
    t.Log("No comm error found")
  }

  if !phoneResponse.Success {
    t.Error("Phone info request should be success")
  }

  if phoneResponse.Type != "landline" {
    t.Error("Phone type should be landline")
  }

  if phoneResponse.Ported {
    t.Error("Phone shouldn't be ported")
  }
}
