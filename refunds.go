package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

//You can perform a partial refund of a Charge.
//The operation can be executed many times, until the entire Charge has been refunded. Once entirely refunded, a Charge can’t be refunded again.
//This method throws an error when called on an already fully refunded Charge, or when trying to refund more money than is left on a it.
//Refund could manage idempotency.
const refundSuffix = "/v1/refunds"

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
