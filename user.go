package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

const userSuffix = "/v1/users"

//CreateUser create a user you want to send Charge request to. The user must be subscribed to satispay service.
//Once you create a user you do not need to create it again but it is enough create a Charge with the user id used previously.
//But don’t worry, if you do not store user id you can call again the Create a user and, for the same phone number,
//it will always return the same user id.
func (client *Client) CreateUser(phoneNumber, idempotencyKey string) (user User, err error) {

	if phoneNumber == "" {
		return User{}, errors.New("Phone number missing")
	}

	bodyStruct := newUser{PhoneNumber: phoneNumber}
	bodyBytes, err := json.Marshal(&bodyStruct)

	if err != nil {
		log.Errorf("Got error marshaling struct %v", err)
		return User{}, err
	}

	request, err := http.NewRequest("POST", client.endpoint+userSuffix, strings.NewReader(string(bodyBytes)))
	if err != nil {
		log.Errorf("Got error during request creation %v", err)
		return User{}, err
	}

	response, err := client.do(request, idempotencyKey)
	if err != nil {
		log.Errorf("Got error in response %v", err)
		return User{}, err
	}

	decoder := json.NewDecoder(response.Body)
	if response.StatusCode == 400 {
		satisErr := SatispayError{}
		err := decoder.Decode(&satisErr)
		if err != nil {
			log.Errorf("Error decoding response body for satispay error %v", err)
			return user, err
		}

		log.Errorf("Got error from api with error code %d and message %s", satisErr.Code, satisErr.Message)
		return user, errors.New(satisErr.Message)
	}
	if response.StatusCode == 404 {
		return User{}, errors.New("The phone number isn’t from a registered user")
	}

	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//UserList get the list of Shop Users of a Online shop.
func (client *Client) UserList(limit int, startingAfter, endingBefore string) ([]User, error) {

	url := composeURL(limit, client.endpoint+userSuffix, startingAfter, endingBefore)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Errorf("Got error building request %v", err)
		return nil, err
	}

	response, err := client.do(request, "")
	if err != nil {
		log.Errorf("Got error perfoming http request %v", err)
		return nil, err
	}

	if response.StatusCode == 400 {
		log.Errorf("Got 400 in user listing")
		return nil, errors.New("Beneficiary validation error")
	}
	if response.StatusCode == 401 {
		log.Errorf("Unauthorized")
		return nil, errors.New("Unauthorized")
	}

	listResp := userListResponse{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&listResp)
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
	url := client.endpoint + userSuffix + userID
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error building request %v", err)
		return User{}, err
	}

	response, err := client.do(request, "")

	if response.StatusCode == 404 {
		return User{}, errors.New("UserShop don’t exist")
	}

	user := User{}
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(user)

	if err != nil {
		log.Errorf("Got error deconding %v", err)
		return user, err
	}

	return user, nil
}
