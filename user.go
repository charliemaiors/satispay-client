package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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

	request, err := http.NewRequest("POST", client.endpoint+"/online/v1/users", strings.NewReader(string(bodyBytes)))
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

}
