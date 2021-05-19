package main

import (
	"testing"
)

type coords struct {
	lat  float64
	long float64
}

func TestDistanceMock(t *testing.T) {
	DelrayBeachFL := coords{lat: 26.45151095322035, long: -80.17595111231378}
	DeerFieldBeachFL := coords{lat: 26.315803607717026, long: -80.0987914408342}
	locator := &MockLocator{}
	min, err := locator.DistanceETAMin(DelrayBeachFL.lat, DelrayBeachFL.long, DeerFieldBeachFL.lat, DeerFieldBeachFL.long)
	if err != nil {
		t.Fatal(err)
	}
	// between these points should take between 8 and 12 minutes
	if min.Minutes() > 12 {
		t.Fatal("distance is over 12 minutes")
	}
	if min.Minutes() < 8 {
		t.Fatal("distance is under 8 minutes")
	}
}
func TestTwilioMock(t *testing.T) {
	service := &MockTwilio{}
	service.SendSMS("Hello world", "5151515")
	recv, err := service.PopLastMessageSent()
	if err != nil {
		t.Fatal(err)
	}
	if recv != "Hello world" {
		t.Fatal("did not recieve correct message back")
	}
}
