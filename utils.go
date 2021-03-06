package client

import (
	"errors"
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//NewChargeRequest create a new ChargeRequest performing validations on parameters as specified in Satispay API documentation https://s3-eu-west-1.amazonaws.com/docs.online.satispay.com/index.html#create-a-charge
func NewChargeRequest(userID, description, currency, callbackURL string, metadata map[string]string, requiredSuccessEmail bool, amount int64, expireIn int) (*ChargeRequest, error) {
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

	if _, err := url.ParseRequestURI(callbackURL); err != nil {
		log.Errorf("Got error validating callback url %v", err)
		return nil, err
	}

	if expireIn < 60 {
		log.Error("Error: expiration time must be at least 60 seconds, using default one (15 minutes, 900 seconds)")
		expireIn = 900
	}

	return &ChargeRequest{Amount: amount, CallBackURL: callbackURL, Currency: currency, Description: description, ExpireIn: expireIn, Metdata: metadata, RequiredSuccessEmail: requiredSuccessEmail, UserID: userID}, nil
}

//NewRefundRequest create a new refund request structure performing validation described in Satispay API https://s3-eu-west-1.amazonaws.com/docs.online.satispay.com/index.html#create-a-refund
func NewRefundRequest(chargeID, description, currency string, amount int64, reason RefundReason, metadata map[string]string) (*RefundRequest, error) {
	if _, err := uuid.FromString(chargeID); err != nil {
		log.Errorf("Invalid charge id %v", err)
		return nil, err
	}

	if len(description) > 255 {
		log.Errorf("Description to long!")
		return nil, errors.New("Description to long")
	}

	if currency != "EUR" {
		log.Errorf("Invalid currency only EUR is supported (until now)")
		return nil, errors.New("Invalid currency only EUR is supported (until now)")
	}

	if amount < 0 {
		log.Errorf("Negative amount")
		return nil, errors.New("Negative amount")
	}

	return &RefundRequest{
		Amount:      amount,
		ChargeID:    chargeID,
		Currency:    currency,
		Description: description,
		Metadata:    metadata,
		Reason:      reason,
	}, nil
}

//NewCheckoutRequest create a new request for checkout
func NewCheckoutRequest(phoneNumber, redirectURL, description, callbackURL, currency string, amount int64) (*CheckoutRequest, error) {
	if phoneNumber == "" {
		log.Errorf("Phone number required")
		return nil, errors.New("Phone number required")
	}

	if amount < 0 {
		log.Errorf("Amount or expiration time negative")
		return nil, errors.New("Amount or expiration time negative")
	}

	if currency != "EUR" {
		log.Errorf("Only EUR currency is supported for the moment")
		return nil, errors.New("Only EUR currency is supported for the moment")
	}

	return &CheckoutRequest{
		AmountUnit:  amount,
		CallbackURL: callbackURL,
		Currency:    currency,
		Description: description,
		PhoneNumber: phoneNumber,
		RedirectURL: redirectURL,
	}, nil
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

func makeTimestamp(date time.Time) int64 {
	return date.UnixNano() / int64(time.Millisecond)
}
