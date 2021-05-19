package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	_ = &MockTwilio{}
	lat1 := flag.Float64("plat", 0.0, "a string")
	long1 := flag.Float64("plong", 0.0, "a string")
	lat2 := flag.Float64("dlat", 0.0, "a string")
	long2 := flag.Float64("dlong", 0.0, "a string")
	order := flag.String("o", "", "a string")
	cellNumber := flag.String("c", "", "Cell number")
	deliveryTime := flag.String("t", "", "in RFC3339 format")

	flag.Parse()

	if deliveryTime == nil || *deliveryTime == "" {
		fmt.Println("Delivery Time Required")
		os.Exit(1)
	}

	// Try and parse the time
	time, err := time.Parse(time.RFC3339, *deliveryTime)
	if err != nil {
		fmt.Println("Date could not be parsed", err)
		os.Exit(1)
	}

	twilio, err := NewTwilioService()
	if err != nil {
		fmt.Println("bad twilio config", err)
		os.Exit(1)
	}
	locator := &MockLocator{}
	srv := NewLateService(twilio, locator)

	err = srv.Handle(&LateServiceRequest{
		customerLat:     lat1,
		customerLong:    long1,
		deliveryLat:     lat2,
		deliveryLong:    long2,
		orderID:         order,
		cellNumber:      cellNumber,
		estDeliveryTime: &time,
	})
	if err != nil {
		fmt.Println("error", err)
	}
}
