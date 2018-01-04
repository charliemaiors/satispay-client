package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

const checkoutsSuffix = "/v1/checkouts"

//CreateCheckout creates a new checkout on Satispay Platform
func (client *Client) CreateCheckout(checkoutRequest *CheckoutRequest, idempotencyKey string) (checkout Checkout, err error) {
	request, err := http.NewRequest("POST", client.endpoint+checkoutsSuffix, strings.NewReader(checkoutRequest.String()))
	if err != nil {
		log.Errorf("Got error creating http request %v", err)
		return checkout, err
	}

	response, err := client.do(request, idempotencyKey)
	if err != nil {
		log.Errorf("Got error performing  http request %v", err)
		return checkout, err
	}

	dec := json.NewDecoder(response.Body)

	if response.StatusCode != 200 {
		satisErr := SatispayError{}
		err = dec.Decode(&satisErr)
		if err != nil {
			log.Errorf("Error decoding satispay error %v", err)
			return checkout, err
		}

		return checkout, errors.New(satisErr.Message)
	}

	err = dec.Decode(&checkout)
	if err != nil {
		log.Errorf("Error deconding checkout %v", err)
	}
	return
}
