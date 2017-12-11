package client

import (
	"net/http"
)

const (
	productionEndpoint = "https://authservices.satispay.com/online"
	sandBoxAPIEndpoint = "https://staging.authservices.satispay.com/online"
)

type Client struct {
	bearerToken string
	endpoint    string
	httpClient  *http.Client
}

type newUser struct {
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

type userListResponse struct {
	HasMore bool   `json:"has_more"`
	List    []User `json:"list"`
}
