package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//You can perform a partial refund of a Charge.
//The operation can be executed many times, until the entire Charge has been refunded. Once entirely refunded, a Charge can’t be refunded again.
//This method throws an error when called on an already fully refunded Charge, or when trying to refund more money than is left on a it.
//Refund could manage idempotency.
const refundSuffix = "/v1/refunds"

//CreateRefund create a refund, you must specify the Charge to create it on.
func (client *Client) CreateRefund(refundRequest *RefundRequest) (Refund, error) {
	url := client.endpoint + refundSuffix
	request, err := http.NewRequest("POST", url, strings.NewReader(refundRequest.String()))
	if err != nil {
		log.Errorf("Got error creating request %v", err)
		return Refund{}, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Got error performing request %v", err)
		return Refund{}, err
	}

	if response.StatusCode == 400 {
		log.Error("Body validation error")
		return Refund{}, errors.New("Body validation error")
	}

	if response.StatusCode == 403 {
		log.Error("Try to create a refund for a Charge not owned by user")
		return Refund{}, errors.New("	Try to create a refund for a Charge not owned by user")
	}

	dec := json.NewDecoder(response.Body)
	refund := Refund{}
	err = dec.Decode(&refund)
	if err != nil {
		log.Errorf("Got error deconding response body %v", err)
		return Refund{}, err
	}

	return refund, nil
}

//GetRefund get a refund by id
func (client *Client) GetRefund(refundID string) (ref Refund, err error) {
	if _, uuidErr := uuid.FromString(refundID); uuidErr != nil {
		log.Errorf("Refund ID is not valid %v", err)
		return ref, uuidErr
	}

	request, err := http.NewRequest("GET", client.endpoint+refundSuffix+"/"+refundID, nil)
	if err != nil {
		log.Errorf("Error creating http request %v", err)
		return
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Error performing http request %v", err)
		return
	}

	if response.StatusCode == 404 {
		log.Errorf("Refund does not exist")
		return ref, errors.New("Refund does not exist")
	}

	if response.StatusCode == 403 {
		log.Errorf("Try to get a refund of another shop")
		return ref, errors.New("Try to get a refund of another shop")
	}

	if response.StatusCode == 400 {
		log.Errorf("Shop validation error")
		return ref, errors.New("Shop validation error")
	}

	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&ref)
	if err != nil {
		log.Errorf("Got error deconding response body %v", err)
	}

	return

}
