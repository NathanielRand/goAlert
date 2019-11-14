package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/subosito/twilio"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	// IEX Cloud API Reference
	// GET https://cloud-sse.iexapis.com/stable/cryptoQuotes?symbols={symbol}&token={YOUR_TOKEN}

	// GET full json response from IEX API forex quotes.
	response, err := http.Get("https://sandbox.iexapis.com/stable/fx/latest?symbols=AUDUSD,USDJPY,EURUSD,USDZAR&token=Tsk_4d1510e26f304e95bcb49028549414a9")
	if err != nil {
		log.Fatalf("Error while calling IEX Cloud API via http.GET %v\n", err)
		return
	}
	// Defer close of http.Get call
	defer response.Body.Close()

	// Grab the entire response body.
	// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended.
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error while reading response body %v\n", err)
		return
	}

	// Print the entire response body to terminal.
	fmt.Println(string(contents))

	// Get the TWILIO_ACCOUNTSID environment variable
	twilioAccountSID, exists := os.LookupEnv("TWILIO_ACCOUNTSID")
	if exists == false {
		log.Fatalf("Error fetching environment variables: %v", exists)
		return
	}

	// Get the TWILIO_AUTH_TOKEN
	twilioAuthToken, exists := os.LookupEnv("TWILIO_AUTH_TOKEN")
	if exists == false {
		log.Fatalf("Error fetching environment variables: %v", exists)
		return
	}

	twilioFromNum, exists := os.LookupEnv("TWILIO_FROM_NUM")
	if exists == false {
		log.Fatalf("Error fetching environment variables: %v", exists)
		return
	}

	twilioToNum, exists := os.LookupEnv("TWILIO_TO_NUM")
	if exists == false {
		log.Fatalf("Error fetching environment variables: %v", exists)
		return
	}

	// Initialize Twilio client
	c := twilio.NewClient(twilioAccountSID, twilioAuthToken, nil)

	// Create message params
	params := twilio.MessageParams{
		Body: string(contents),
	}

	// Send SMS via Twilio
	s, resp, err := c.Messages.Send(twilioFromNum, twilioToNum, params)
	log.Println("Send:", s)
	log.Println("Response:", resp)
	log.Println("Err:", err)
}
