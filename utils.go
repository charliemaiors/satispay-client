package client

import (
	"errors"
	"net/url"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//NewChargeRequest create a new ChargeRequest performing validations on parameters as specified in Satispay API documentation
//https://s3-eu-west-1.amazonaws.com/docs.online.satispay.com/index.html#create-a-charge
func NewChargeRequest(userID, description, currency, callbackUrl string, metadata map[string]string, requiredSuccessEmail bool, amount int64, expireIn int) (*ChargeRequest, error) {
	_, err := uuid.FromString(userID)

	if err != nil {
		log.Errorf("UserID is not a valid UUID: %v", err)
		return nil, err
	}

	if currency != "EUR" { //Temporary check, only EURO is supported for currency
		log.Error("Only EUR is supported for currency by Satispay Platform")
		return nil, errors.New("Only EUR is supported for currency by Satispay Platform")
	}

	if expireIn < 0 {
		log.Error("Invalid value for expiration time")
		return nil, errors.New("Invalid value for expiration time")
	}

	if amount < 0 {
		log.Error("Invalid value for amount")
		return nil, errors.New("Invalid value for amount")
	}

	if _, err := url.ParseRequestURI(callbackUrl); err != nil {
		log.Errorf("Got error validating callback url %v", err)
		return nil, err
	}

	if expireIn < 60 {
		log.Error("Error: expiration time must be at least 60 seconds, using default one (15 minutes, 900 seconds)")
		expireIn = 900
	}

	return &ChargeRequest{Amount: amount, CallBackURL: callbackUrl, Currency: currency, Description: description, ExpireIn: expireIn, Metdata: metadata, RequiredSuccessEmail: requiredSuccessEmail, UserID: userID}, nil
}

func NewRefundRequest() {
	
}

func composeURL(limit int, initialURL, startingAfter, endingBefore string) string {
	url := initialURL + "?"
	if limit > 0 && limit < 100 {
		url += "limit=" + string(limit)
	} else {
		url += "limit=20"
	}
	if startingAfter != "" {
		_, err := uuid.FromString(startingAfter)

		if err != nil {
			log.Errorf("Invalid initial UUID, parsing error: %v\nIgnoring parameter", err)
		} else {
			url += "&starting_after=" + startingAfter
		}
	}
	if endingBefore != "" {
		_, err := uuid.FromString(endingBefore)

		if err != nil {
			log.Errorf("Invalid ending UUID, parsing error: %v\nIgnoring parameter", err)
		} else {
			url += "&ending_before=" + endingBefore
		}
	}
	return url
}
