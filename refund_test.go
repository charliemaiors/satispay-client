package client_test

import (
	"testing"

	client "github.com/charliemaiors/satispay-client"
)

func TestNewRefundUUID(test *testing.T) {
	_, err := client.NewRefundRequest("dhofoadisx3j.", "Test charge request", "EU", -1, client.Fraudolent, nil)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestNewRefundCurrency(test *testing.T) {
	_, err := client.NewRefundRequest("68170747-ae17-4799-9698-9059b550f2f0", "Test charge request", "EU", -1, client.Fraudolent, nil)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestNewRefundAmount(test *testing.T) {
	_, err := client.NewRefundRequest("68170747-ae17-4799-9698-9059b550f2f0", "Test charge request", "EU", -1, client.Fraudolent, nil)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientGetRefund(test *testing.T) {
	_, err := satisClient.GetRefund("sdhuishufdihfsa.jdi")
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientGetRefundList(test *testing.T) {
	_, err := satisClient.GetRefundList(40, "sdhuishufdihfsa.jdi", "fidahufihsfiushafui-", "dk30-2-2-d")
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}

func TestClientUpdateRefundFail(test *testing.T) {
	_, err := satisClient.UpdateRefund("fsdhiahfusd-d", nil)
	if err == nil {
		test.Fatal("Expecting error creating new request")
	}
	test.Logf("Error is %v", err)
}
