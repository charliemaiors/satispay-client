package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

const chargeSuffix = "/v1/charges"

//CreateCharge create a Charge having a user id.
func (client *Client) CreateCharge(chargeRequest *ChargeRequest) (Charge, error) {

	reader := strings.NewReader(chargeRequest.String())
	request, err := http.NewRequest("POST", client.endpoint+chargeSuffix, reader)

	if err != nil {
		log.Errorf("Got error creating the request %v", err)
		return Charge{}, err
	}

	response, err := client.do(request)

	if err != nil {
		log.Errorf("Got error performing the request %v", err)
		return Charge{}, err
	}

	if response.StatusCode == 400 {
		log.Error("Got 400: body validation error")
		return Charge{}, errors.New("Body validation error")
	}

	if response.StatusCode == 403 {
		log.Error("Got 403: Trying to create a Charge for another user")
		return Charge{}, errors.New("Trying to create a Charge for another user")
	}

	dec := json.NewDecoder(response.Body)
	charge := Charge{}
	err = dec.Decode(&charge)

	if err != nil {
		log.Errorf("Got error deconding body %v", err)
		return Charge{}, err
	}

	return charge, nil
}

//GetChargeList get a list of Charge ordered by creation.
//To get element staring after or ending before (excluse) a Charge passed populate starting_after or ending_before with the Charge id.
//If both starting_after and ending before elements are passed return element ending before the id passed.
//Limit value indicate the number of elements returned.
//You could also pass a starting_after_timestamp query param with a UNIX timestamp in mills, will be returned the Charges after that date.
func (client *Client) GetChargeList(limit int, startingAfter, endingBefore string, startingAfterDate time.Time) ([]Charge, error) {
	if _, err := uuid.FromString(startingAfter); err != nil {
		log.Errorf("Starting after is not valid uuid %v", err)
		return nil, err
	}

	if _, err := uuid.FromString(endingBefore); err != nil {
		log.Errorf("Ending before is not valid uuid %v", err)
		return nil, err
	}

	if limit > 100 || limit < 0 {
		return nil, errors.New("Invalid limit value")
	}

	url := composeURL(limit, client.endpoint+chargeSuffix, startingAfter, endingBefore)
	if !startingAfterDate.IsZero() {
		if startingAfterDate.After(time.Now()) {
			return nil, errors.New("Invalid date, is after today")
		}
		url += "&starting_after_timestamp=" + string(makeTimestamp(startingAfterDate))
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error building request %v", err)
		return nil, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Error performing http request %v", err)
		return nil, err
	}

	if response.StatusCode == 403 {
		log.Errorf("Try to get a Charge of another shop")
		return nil, errors.New("Try to get a Charge of another shop")
	}

	if response.StatusCode == 400 {
		log.Errorf("Beneficiary validation")
		return nil, errors.New("Beneficiary validation")
	}

	listResult := chargeListResponse{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&listResult)
	if err != nil {
		log.Errorf("Got error decoding response body %v", err)
		return nil, err
	}

	return listResult.List, nil
}

//GetCharge get a Charge by id
func (client *Client) GetCharge(chargeID string) (Charge, error) {
	if _, err := uuid.FromString(chargeID); err != nil {
		log.Errorf("Invalid charge ID %v", err)
		return Charge{}, err
	}

	url := client.endpoint + chargeSuffix + "/" + chargeID
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error generating http request %v", err)
		return Charge{}, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Error performing http request %v", err)
		return Charge{}, err
	}

	if response.StatusCode == 404 {
		log.Errorf("Charge does not exist")
		return Charge{}, errors.New("Charge does not exist")
	}

	if response.StatusCode == 403 {
		log.Errorf("Try to get a Charge of another shop")
		return Charge{}, errors.New("Try to get a Charge of another shop")
	}

	if response.StatusCode == 400 {
		log.Errorf("Shop validation error")
		return Charge{}, errors.New("Shop validation error")
	}

	charge := Charge{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&charge)
	if err != nil {
		log.Errorf("Error deconding http response body %v", err)
		return Charge{}, err
	}

	return charge, nil
}

//UpdateCharge update a Charge, only metadata, description and state can be updated.
//About parameters:
//Metadata - Object key-value - not mandatory - Object with max 20 keys. Key length 45 characters and values 500 characters. New keys will be added, the existing keys will be updated and keys set with null value will be removed.
//Description -	string - not mandatory - a Charge description
//ChargeState - bool -	not mandatory - if set the target Charge gets canceled; the staus will be set to FAILURE and the status_detail will be set to DECLINED_BY_PAYER.
func (client *Client) UpdateCharge(chargeID, description string, metadata map[string]string, changeState bool) (Charge, error) {
	if _, err := uuid.FromString(chargeID); err != nil {
		log.Errorf("Invalid charge UUID %v", err)
		return Charge{}, err
	}

	updateStruct := chargeUpdate{Description: description, Metadata: metadata}

	if changeState {
		updateStruct.ChangeState = "CANCELED"
	}

	url := client.endpoint + chargeSuffix + "/" + chargeID
	request, err := http.NewRequest("PUT", url, strings.NewReader(updateStruct.String()))
	if err != nil {
		log.Errorf("Got error creating request %v", err)
		return Charge{}, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Got error performing request %v", err)
		return Charge{}, err
	}

	if response.StatusCode == 403 {
		log.Error("Try to update a Charge of another user or Try to cancel a Charge which is already in state SUCCESS")
		return Charge{}, errors.New("Try to update a Charge of another user or Try to cancel a Charge which is already in state SUCCESS")
	}

	if response.StatusCode == 400 {
		log.Error("Beneficiary validation or body validation error")
		return Charge{}, errors.New("Beneficiary validation or body validation error")
	}

	if response.StatusCode == 404 {
		log.Error("Charge does not exist")
		return Charge{}, errors.New("Charge does not exist")
	}

	charge := Charge{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&charge)
	if err != nil {
		log.Errorf("Got error deconding response body %v", err)
		return Charge{}, err
	}

	return charge, nil
}
