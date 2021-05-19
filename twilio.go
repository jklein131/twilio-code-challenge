package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type TwilioService struct {
	accountSID string
	authToken  string
	fromNum    string
	client     *http.Client
}

var _ SMSServicer = &TwilioService{} // or &myType{} or [&]myType if scalar

func NewTwilioService() (*TwilioService, error) {
	twilioAccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	twilioFromNumber := os.Getenv("TWILIO_FROM_NUMBER")
	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if twilioAccountSID == "" {
		return nil, errors.New("Enviroment variable TWILIO_ACCOUNT_SID must be set")
	}
	if twilioAuthToken == "" {
		return nil, errors.New("Enviroment variable TWILIO_AUTH_TOKEN must be set")
	}
	if twilioFromNumber == "" {
		return nil, errors.New("Enviroment variable TWILIO_FROM_NUMBER must be set")
	}
	return &TwilioService{
		accountSID: twilioAccountSID,
		authToken:  twilioAuthToken,
		fromNum:    twilioFromNumber,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

func (m *TwilioService) SendSMS(s string, to string) error {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", m.fromNum)
	msgData.Set("Body", s)

	msgDataReader := *strings.NewReader(msgData.Encode())
	req, _ := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+m.accountSID+"/Messages.json", &msgDataReader)
	req.SetBasicAuth(m.accountSID, m.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	} else {
		t, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(t))
	}
}
