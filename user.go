package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func (client *Client) CreateUser(phoneNumber string) (User, error) {

	if phoneNumber == "" {
		return User{}, errors.New("Phone number missing")
	}

	bodyStruct := newUser{PhoneNumber: phoneNumber}
	bodyBytes, err := json.Marshal(&bodyStruct)

	if err != nil {
		log.Errorf("Got error marshaling struct %v", err)
		return User{}, err
	}

	request, err := http.NewRequest("POST", client.endpoint+"/v1/users", strings.NewReader(string(bodyBytes)))
	if err != nil {
		log.Errorf("Got error during request creation %v", err)
		return User{}, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Got error in response %v", err)
		return User{}, err
	}

	if response.StatusCode == 400 {
		return User{}, errors.New("Something bad happened, could be an invalid phone number or shop validation error")
	}
	if response.StatusCode == 404 {
		return User{}, errors.New("The phone number isn’t from a registered user")
	}

	decoder := json.NewDecoder(response.Body)
	user := User{}
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (client *Client) UserList(limit int, startingAfter, endingBefore string) ([]User, error) {

	url := composeURL(limit, client.endpoint, startingAfter, endingBefore)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Errorf("Got error building request %v", err)
		return nil, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Got error perfoming http request %v", err)
		return nil, err
	}

	if response.StatusCode == 400 {
		log.Errorf("Got 400 in user listing")
		return nil, errors.New("Beneficiary validation error")
	}

	listResp := userListResponse{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(listResp)
	if err != nil {
		log.Errorf("Got error deconding %v", err)
		return nil, err
	}
	log.Debugf("Has more? %v", listResp.HasMore)
	return listResp.List, nil
}

func (client *Client) GetUser(userID string) (User, error) {
	if _, err := uuid.FromString(userID); err != nil {
		return User{}, err
	}
	url := client.endpoint + "/v1/users/" + userID
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error building request %v", err)
		return User{}, err
	}

	response, err := client.do(request)

	user := User{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(user)

	if err != nil {
		log.Errorf("Got error deconding %v", err)
		return user, err
	}

	return user, nil
}

func composeURL(limit int, initialURL, startingAfter, endingBefore string) string {
	url := initialURL + "/v1/users?"
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
