package client_test

import (
	"testing"
	"time"

	client "github.com/charliemaiors/satispay-client"
)

var satisClient *client.Client
var err error

func init() {
	satisClient, err = client.NewClient("sdhofdhafshduiahufdafhsudodfshauoo", false)
	if err != nil {
		panic(err)
	}
}

func TestNewChargeUUID(test *testing.T) {
	_, err := client.NewChargeRequest("dhofoadisx3j.", "Test charge request", "USD", "http://ciaone.org/", nil, false, 300, -1)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestNewChargeCurrency(test *testing.T) {
	_, err := client.NewChargeRequest("68170747-ae17-4799-9698-9059b550f2f0", "Test charge request", "USD", "http://ciaone.org/", nil, false, 300, -1)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestNewChargeExpireIn(test *testing.T) {
	_, err := client.NewChargeRequest("68170747-ae17-4799-9698-9059b550f2f0", "Test charge request", "EUR", "http://ciaone.org/", nil, false, 300, -1)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientGetCharge(test *testing.T) {
	_, err := satisClient.GetCharge("sdhuishufdihfsa.jdi")
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientGetCharges(test *testing.T) {
	_, err := satisClient.GetChargeList(40, "sdhuishufdihfsa.jdi", "fidahufihsfiushafui-", time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientUpdateChargeFail(test *testing.T) {
	_, err := satisClient.UpdateCharge("fsdhiahfusd-d", "Test update fake charge", nil, true)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}
