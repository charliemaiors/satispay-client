package client

import (
	"net/http"
	"time"
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

type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

type ChargeRequest struct {
	UserID               string            `json:"user_id"`
	Description          string            `json:"description"`
	Currency             string            `json:"currency"`
	CallBackURL          string            `json:"callback_url"`
	Amount               int64             `json:"amount"`
	Metdata              map[string]string `json:"metadata"`
	RequiredSuccessEmail bool              `json:"required_success_email"`
	ExpireIn             int               `json:"expire_in"`
}

type Charge struct {
	ID                   string            `json:"id"`
	Description          string            `json:"description"`
	Currency             string            `json:"currency"`
	Amount               int64             `json:"amount"`
	Status               ChargeStatus      `json:"status"`
	UserID               string            `json:"user_id"`
	UserShortName        string            `json:"user_short_name"`
	Metadata             map[string]string `json:"metadata"`
	RequiredSuccessEmail bool              `json:"required_success_email"`
	ExpireDate           time.Time         `json:"expire_date"`
	CallbackURL          string            `json:"callback_url"`
}

// Private types and constants

const (
	productionEndpoint = "https://authservices.satispay.com/online"
	sandBoxAPIEndpoint = "https://staging.authservices.satispay.com/online"
)

type newUser struct {
	PhoneNumber string `json:"phone_number"`
}

type userListResponse struct {
	HasMore bool   `json:"has_more"`
	List    []User `json:"list"`
}
