package main

import "time"

type SMSServicer interface {
	SendSMS(msg string, to string) error
}

type LocationServicer interface {
	DistanceETAMin(lat1, long1, lat2, long2 float64) (time.Duration, error)
}
