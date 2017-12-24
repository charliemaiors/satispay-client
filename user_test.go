package client_test

import (
	"testing"
)

func TestNewUserNoPhoneNumber(test *testing.T) {
	_, err := satisClient.CreateUser("")
	if err == nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}
	test.Logf("The beautifull error message is %v", err)
}

func TestGetUser(test *testing.T) {
	_, err := satisClient.GetUser("fjoiafdaks")
	if err == nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}
	test.Logf("The beautifull error message is %v", err)
}

func TestGetUserList(test *testing.T) {
	_, err := satisClient.UserList(40, "fdhuaihfaos", "fghufhifsdahiaf")
	if err == nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}
	test.Logf("The beautifull error message is %v", err)
}
