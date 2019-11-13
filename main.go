package main

import (
	"log"

	"github.com/subosito/twilio"
)

var (
	AccountSid = "AC35ddabb983353c74dc8f9149db26a5f5"
	AuthToken  = "bacba053badfeb5ecce41afcbef528fa"
	From       = "+18437906751"
	To         = "+18434696850"
)

func main() {
	// Initialize Twilio client
	c := twilio.NewClient(AccountSid, AuthToken, nil)

	// Send SMS
	// Add FX info to body
	params := twilio.MessageParams{
		Body: "goAlertFX",
	}

	s, resp, err := c.Messages.Send(From, To, params)
	log.Println("Send:", s)
	log.Println("Response:", resp)
	log.Println("Err:", err)
}
