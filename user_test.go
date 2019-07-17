package client_test

import (
	"os"
	"testing"
)

func TestNewUserNoPhoneNumber(test *testing.T) {
	_, err := satisClient.CreateUser("", "")
	if err == nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}
	test.Logf("The beautifull error message is %v", err)
}

func TestNewUserInvalid(test *testing.T) {
	_, err := validSatisClient.CreateUser("+393210987654", "")
	if err == nil {
		test.Fatalf("User creation has accomplished instead of expected failure")
	}
	test.Logf("The error is %v", err)
}
func TestNewUserValid(test *testing.T) {
	user, err := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "")
	if err != nil {
		test.Fatalf("User creation has failed %v", err)
	}
	test.Logf("User id is %s with phone number %s", user.ID, user.PhoneNumber)
}

func TestNewUserValidIdempotency(test *testing.T) {
	user, err := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "1")
	if err != nil {
		test.Fatalf("User creation has failed %v", err)
	}
	test.Logf("User id is %s with phone number %s", user.ID, user.PhoneNumber)

	user2, err := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "1")
	if err != nil {
		test.Fatalf("User2 creation has failed %v", err)
	}
	test.Logf("User id is %s with phone number %s", user2.ID, user2.PhoneNumber)

	if user != user2 {
		test.Fatal("User and User2 are not equals")
	}

	test.Logf("User and User2 are equal? %v", user == user2)
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
		test.Fatalf("Expecting error nil but instead is %v", err)
	}

	test.Logf("The beautifull error message is %v", err)
}

func TestGetUserListValid(test *testing.T) {
	list, err := validSatisClient.UserList(40, "", "")
	if err != nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}

	test.Logf("The user list is %v", list)
}
