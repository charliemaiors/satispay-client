package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

//NewClient define a new instance of Client struct, at creation time it will ONLY check if token is not equals
//to an empty string. If is empty the function will return an error otherwise a pointer to the built client
func NewClient(token string, isProduction bool) (*Client, error) {

	if token == "" {
		return nil, errors.New("Invalid token")
	}

	client := &Client{bearerToken: token}

	if isProduction {
		client.endpoint = productionEndpoint
	} else {
		client.endpoint = sandBoxAPIEndpoint
	}

	client.httpClient = http.DefaultClient
	return client, nil
}

//CheckBearer refers to https://s3-eu-west-1.amazonaws.com/docs.online.satispay.com/index.html#api-check-bearer api, this method will check throught satispay api
//if the provided token is valid, is highly recommended to use this method after client creation at least the first time
func (client *Client) CheckBearer() bool {
	request, err := http.NewRequest("GET", client.endpoint+"/wally-services/protocol/authenticated", nil)

	if err != nil {
		log.Errorf("Got error creating request %v", err)
		return false
	}

	response, err := client.do(request)

	if err != nil {
		log.Errorf("Got error in request %v", err)
	}

	if response.StatusCode == 401 || response.StatusCode == 403 {
		return false
	}

	return true
}

func (client *Client) CreateUser(phoneNumber string) (*User, error) {

	if phoneNumber == "" {
		return nil, errors.New("Phone number missing")
	}

	bodyStruct := newUser{PhoneNumber: phoneNumber}
	bodyBytes, err := json.Marshal(&bodyStruct)

	if err != nil {
		log.Errorf("Got error marshaling struct %v", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", client.endpoint+"/online/v1/users", strings.NewReader(string(bodyBytes)))
	if err != nil {
		log.Errorf("Got error during request creation %v", err)
		return nil, err
	}

	response, err := client.do(request)
	if err != nil {
		log.Errorf("Got error in response %v", err)
		return nil, err
	}

	if response.StatusCode == 400 {
		return nil, errors.New("Something bad happened, could be an invalid phone number or shop validation error")
	}
	if response.StatusCode == 404 {
		return nil, errors.New("The phone number isn’t from a registered user")
	}

	decoder := json.NewDecoder(response.Body)
	user := &User{}
	err = decoder.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) do(request *http.Request) (*http.Response, error) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+client.bearerToken)
	return client.httpClient.Do(request)
}
