package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const amountSuffix = "/v1/amounts"

//GetTotalAmount calculate the total amount of Charges in SUCCESS status and the total amount of Refunds within a specific timeframe.
func (client *Client) GetTotalAmount(startingDate, endingDate time.Time) (amount TotalAmount, err error) {
	if endingDate.Before(startingDate) {
		return amount, errors.New("Ending date is before starting date")
	}

	if startingDate.IsZero() || endingDate.IsZero() {
		log.Errorf("")
		return amount, errors.New("Starting date or ending date is zero, please provide a valid date")
	}

	url := client.endpoint + amountSuffix + "?starting_date=" + string(makeTimestamp(startingDate)) + "&ending_date=" + string(makeTimestamp(endingDate))
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Got error creating the request %v", err)
		return amount, err
	}

	response, err := client.do(request, "")
	if err != nil {
		log.Errorf("Got error performing the http request, details: %v", err)
		return
	}

	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&amount)
	if err != nil {
		log.Errorf("Got error deconding body %v", err)
	}

	return
}
