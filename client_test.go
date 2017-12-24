package client_test

import (
	"testing"

	client "bitbucket.org/cmaiorano/satispay-client"
)

func TestNewClientError(test *testing.T) {
	_, err := client.NewClient("", false)
	if err == nil {
		test.Fatal("Error is not nil but token is empty, test failed")
	}
	test.Logf("Error is %v", err)
}

func TestClientBearer(test *testing.T) {
	client, err := client.NewClient("dfsfsdafdaadfs", false)
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}

	valid := client.CheckBearer()

	if valid {
		test.Fatalf("Expecting random generated token invalid but instead is %v", valid)
	}
	test.Logf("Token is valid %v", valid)
}
