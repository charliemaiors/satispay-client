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
