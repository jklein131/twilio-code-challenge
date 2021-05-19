package main

import (
	"errors"
	"fmt"
	"time"
)

type LateService struct {
	smsService      SMSServicer
	locationService LocationServicer
}
type LateServiceRequest struct {
	customerLat     *float64
	customerLong    *float64
	deliveryLat     *float64
	deliveryLong    *float64
	orderID         *string
	cellNumber      *string
	orderTime       *time.Time
	estDeliveryTime *time.Time
}

func NewLateService(sms SMSServicer, location LocationServicer) *LateService {
	return &LateService{
		smsService:      sms,
		locationService: location,
	}
}
func (s *LateService) checkReq(req *LateServiceRequest) error {
	if req.orderID == nil || *req.orderID == "" {
		return errors.New("Order ID Required")
	}
	// if req.orderTime == nil || !req.estDeliveryTime.After(*req.orderTime) {
	// 	return errors.New("Delivery Time Must Be After Order Time")
	// }
	if req.deliveryLat == nil ||
		req.deliveryLong == nil {
		return errors.New("Delivery Lat/Long Required")
	}
	if req.customerLat == nil ||
		req.customerLong == nil {
		return errors.New("Customer Lat/Long Required")
	}
	// TODO: add phone number validation here using libphonenumber
	return nil
}

func (s *LateService) Handle(req *LateServiceRequest) error {
	if err := s.checkReq(req); err != nil {
		return err
	}
	dMin, err := s.locationService.DistanceETAMin(*req.customerLat, *req.customerLong, *req.deliveryLat, *req.deliveryLong)
	if err != nil {
		return err
	}

	if time.Now().Add(dMin).After((*req.estDeliveryTime)) {
		msg := fmt.Sprintf("Your delivery %v will be late by %.0f min\n", *req.orderID, time.Now().Add(dMin).Sub((*req.estDeliveryTime)).Minutes())
		err := s.smsService.SendSMS(msg, *req.cellNumber)
		if err != nil {
			return err
		}
		fmt.Println("Send SMS", msg)
		return nil
	}
	fmt.Println("No SMS to send")
	return nil
}
