package client_test

import (
	"os"
	"testing"

	client "github.com/charliemaiors/satispay-client"
	log "github.com/sirupsen/logrus"
)

var validSatisClient, satisClient *client.Client
var err error

func init() {
	if token := os.Getenv("BEARER_TOKEN"); token == "" {
		panic("In order to run tests please set environment variable BEARER_TOKEN with valid satispay token")
	} else {
		validSatisClient, err = client.NewClient(token, false)
		if err != nil {
			panic(err)
		}
	}

	satisClient, err = client.NewClient("sdhofdhafshduiahufdafhsudodfshauoo", false)
	if err != nil {
		panic(err)
	}
	log.SetLevel(log.DebugLevel)
}

func TestNewClientError(test *testing.T) {
	_, err := client.NewClient("", false)
	if err == nil {
		test.Fatal("Error is not nil but token is empty, test failed")
	}
	test.Logf("Error is %v", err)
}

func TestClientBearer(test *testing.T) {

	valid := satisClient.CheckBearer()

	if valid {
		test.Fatalf("Expecting random generated token invalid but instead is %v", valid)
	}
	test.Logf("Token is valid %v", valid)
}

func TestClientBearerValid(test *testing.T) {

	valid := validSatisClient.CheckBearer()

	if !valid {
		test.Fatalf("Expecting random generated token invalid but instead is %v", valid)
	}
	test.Logf("Token is valid %v", valid)
}
