// Code generated by "stringer -type=RefundReason"; DO NOT EDIT.

package client

import "fmt"

const _RefundReason_name = "DuplicateFraudolentRequestedByCustomer"

var _RefundReason_index = [...]uint8{0, 9, 19, 38}

func (i RefundReason) String() string {
	if i < 0 || i >= RefundReason(len(_RefundReason_index)-1) {
		return fmt.Sprintf("RefundReason(%d)", i)
	}
	return _RefundReason_name[_RefundReason_index[i]:_RefundReason_index[i+1]]
}