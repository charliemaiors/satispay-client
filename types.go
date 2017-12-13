package client

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
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

//RefundReason indicating the reason for the refund.
type RefundReason int

const (
	//Duplicate means a charge payed twice for some reason
	Duplicate RefundReason = iota

	//Fraudolent means that a charge is fraudolent
	Fraudolent

	//RequestedByCustomer for other reason requested by customer
	RequestedByCustomer
)

//Client is the main structure of this library, represent the main client in order to interact with Satispay platform
type Client struct {
	bearerToken string
	endpoint    string
	httpClient  *http.Client
}

//User represent a Satispay user resource
type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
}

//ChargeRequest represent a Satispay charge request for a target user identified by its id
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

//Charge represent a Satispay charge
type Charge struct {
	ID                   string             `json:"id"`
	Description          string             `json:"description"`
	Currency             string             `json:"currency"`
	Amount               int64              `json:"amount"`
	Status               ChargeStatus       `json:"status"`
	StatusDetail         ChargeStatusDetail `json:"status_detail"`
	UserID               string             `json:"user_id"`
	UserShortName        string             `json:"user_short_name"`
	Metadata             map[string]string  `json:"metadata"`
	RequiredSuccessEmail bool               `json:"required_success_email"`
	ExpireDate           time.Time          `json:"expire_date"`
	CallbackURL          string             `json:"callback_url"`
}

//RefundRequest represent the request for refund in Satispay Platform
type RefundRequest struct {
	ChargeID    string            `json"charge_id"`
	Description string            `json"description"`
	Amount      int64             `json"amount"`
	Currency    string            `json"currency"`
	Reason      RefundReason      `json"reason"`
	Metadata    map[string]string `json"metadata"`
}

type Refund struct {
	ID string `json"id"`
	*RefundRequest
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

type chargeListResponse struct {
	HasMore bool     `json:"has_more"`
	List    []Charge `json:"list"`
}

type chargeUpdate struct {
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	ChangeState string            `json:"change_state,omitempty`
}

//String is the implementation of Stringer interface for ChargeRequest
func (request *ChargeRequest) String() string {
	jsonifiedRequest, err := json.Marshal(request)

	if err != nil {
		log.Errorf("Got error while marshaling request %v", err)
		return ""
	}

	return string(jsonifiedRequest)
}

func (update chargeUpdate) String() string {
	jsonifiedUpdate, err := json.Marshal(update)

	if err != nil {
		log.Errorf("Got error while marshaling update %v", err)
		return ""
	}

	return string(jsonifiedUpdate)
}
