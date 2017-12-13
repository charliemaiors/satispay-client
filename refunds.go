package client

//You can perform a partial refund of a Charge.
//The operation can be executed many times, until the entire Charge has been refunded. Once entirely refunded, a Charge can’t be refunded again.
//This method throws an error when called on an already fully refunded Charge, or when trying to refund more money than is left on a it.
//Refund could manage idempotency.
const refundSuffix = "/v1/refunds"

func (client *Client) CreateRefund() {

}
