package client

import (
	"net/http"
)

const (
	productionEndpoint = "https://authservices.satispay.com/online"
	sandBoxAPIEndpoint = "https://staging.authservices.satispay.com/online"
)

//ChargeStatus represent the status of requested Charge
type ChargeStatus int

const (
	//Required Charge sent to a user waitng for acceptance
	Required ChargeStatus = iota

	//Success Charge accepted by the user
	Success

	//Failure Charge failed, more details can be found on ChargeStatusDetail
	Failure
)

//ChargeStatusDetail represent the detail regarding a failure of Charge operation
type ChargeStatusDetail int

const (
	//DeclinedByPayer user declined the Charge
	DeclinedByPayer ChargeStatusDetail = iota

	//DeclinedByPayerNotRequired user declined the Charge because he did not request it
	DeclinedByPayerNotRequired

	//CancelByNewCharge same Charge sent to the same user, the second will override the first
	CancelByNewCharge

	//InternalFailure generic error
	InternalFailure

	//Expired the Charge has expired
	Expired
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
