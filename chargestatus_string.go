package client

import "fmt"

const _ChargeStatus_name = "REQUIREDSUCCESSFAILURE"

var _ChargeStatus_index = [...]uint8{0, 8, 15, 22}

func (i ChargeStatus) String() string {
	if i < 0 || i >= ChargeStatus(len(_ChargeStatus_index)-1) {
		return fmt.Sprintf("ChargeStatus(%d)", i)
	}
	return _ChargeStatus_name[_ChargeStatus_index[i]:_ChargeStatus_index[i+1]]
}
