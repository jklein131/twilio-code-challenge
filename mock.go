package main

import (
	"errors"
	"time"
)

type MockLocator struct{}

var _ LocationServicer = &MockLocator{} // or &myType{} or [&]myType if scalar

func (m *MockLocator) DistanceETAMin(lat1, long1, lat2, long2 float64) (time.Duration, error) {

	// To calculate the distance in the straight line, we can use the Haversine function
	d := Haversine(lat1, long1, lat2, long2)
	// THe distance is minutes is based on a speed of 1 mile per minute.
	// Multiplying our duration by 60 gives us the duration it would take to go that distance.
	return time.Second * time.Duration(d*60), nil
}

type MockTwilio struct {
	messages []string
}

var _ SMSServicer = &MockTwilio{} // or &myType{} or [&]myType if scalar

func (m *MockTwilio) SendSMS(s string, to string) error {
	m.messages = append(m.messages, s)
	return nil
}

func (m *MockTwilio) PopLastMessageSent() (string, error) {
	if len(m.messages) == 0 {
		return "", errors.New("no message sent")
	}
	lastMessage := m.messages[len(m.messages)-1]
	m.messages = m.messages[:len(m.messages)-1]
	return lastMessage, nil
}
