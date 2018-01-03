package client_test

import (
	"testing"
	"time"
)

func TestWrongAmountDates(test *testing.T) {
	_, err := satisClient.GetTotalAmount(time.Now(), time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	if err == nil {
		test.Fatalf("Expecting error not nil but instead is %v", err)
	}
	test.Logf("The beautifull error is %v", err)
}

func TestValidAmount(test *testing.T) {
	amounts, err := validSatisClient.GetTotalAmount(time.Date(2017, time.December, 27, 0, 0, 0, 0, time.UTC), time.Now())
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}
	test.Logf("Amounts are %v", amounts)
}
