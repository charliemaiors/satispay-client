package client_test

/*
func TestNewCheckout(test *testing.T) {
	user, err := validSatisClient.CreateUser(os.Getenv("PHONE_NUMBER"), "1")
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}

	chargeReq, err := client.NewChargeRequest(user.ID, "Test", "EUR", "http://test.org", nil, false, 6400, 200)
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}

	charge, err := validSatisClient.CreateCharge(chargeReq, "2")
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}
	test.Logf("Charge is %s", charge.String())

	req, err := client.NewCheckoutRequest(os.Getenv("PHONE_NUMBER"), "http://test.org", "", "http://test.org/mycallback?charge="+charge.ID, "EUR", 6400)
	if err != nil {
		test.Fatalf("Expecting error nil but instead is %v", err)
	}

	checkout, err := validSatisClient.CreateCheckout(req, "1")
	if err != nil {
		test.Fatalf("Expecting checkout error nil but instead is %v", err)
	}
	test.Logf("Checkout is %v", checkout)
}
*/
