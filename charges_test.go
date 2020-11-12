package client_test

import (
	"testing"
	"time"

	client "github.com/charliemaiors/satispay-client"
)

func TestNewChargeUUID(test *testing.T) {
	_, err := client.NewChargeRequest("dhofoadisx3j.", "Test charge request", "USD", "http://ciaone.org/", nil, false, 300, -1)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

/*
func TestNewChargeValidUUID(test *testing.T) {
	user, _ := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "")
	chargeReq, err := client.NewChargeRequest(user.ID, "Test charge request", "EUR", "http://ciaone.org/", nil, false, 300, 200)
	if err != nil {
		test.Fatalf("Expecting error nil creating new request, but instead is %v", err)
	}

	charge, err := validSatisClient.CreateCharge(chargeReq, "")
	if err != nil {
		test.Fatalf("Expecting error nil performing new request, but instead is %v", err)
	}
	test.Logf("Charge is %+v", charge)

	updatedCharge, err := validSatisClient.UpdateCharge(charge.ID, "Valid update", nil, true)
	if err != nil {
		test.Fatalf("Expecting error nil updating request, but instead is %v", err)
	}

	test.Logf("Updated request is %v", updatedCharge)
}

func TestNewChargeValidUUIDIdempotent(test *testing.T) {
	user, _ := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "")
	chargeReq, err := client.NewChargeRequest(user.ID, "Test charge request", "EUR", "http://ciaone.org/", nil, false, 300, 60)
	if err != nil {
		test.Fatalf("Expecting error nil creating new request, but instead is %v", err)
	}

	charge, err := validSatisClient.CreateCharge(chargeReq, "1")
	if err != nil {
		test.Fatalf("Expecting error nil performing new request, but instead is %v", err)
	}
	test.Logf("Charge is %+v", charge)

	charge2, err := validSatisClient.CreateCharge(chargeReq, "1")
	if err != nil {
		test.Fatalf("Expecting error nil performing new request, but instead is %v", err)
	}
	test.Logf("Charge is %+v", charge2)

	if charge.ID != charge2.ID {
		test.Fatal("Expected equals charge structure but instead are different")
	}
	test.Logf("Charges are equal? %v", charge.ID == charge2.ID)
}
*/

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
