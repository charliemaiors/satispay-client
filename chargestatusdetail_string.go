package client

import "fmt"

const _ChargeStatusDetail_name = "DECLINED_BY_PAYERDECLINED_BY_PAYER_NOT_REQUIREDCANCEL_BY_NEW_CHARGEINTERNAL_FAILUREEXPIRED"

var _ChargeStatusDetail_index = [...]uint8{0, 17, 47, 67, 83, 90}

func (i ChargeStatusDetail) String() string {
	if i < 0 || i >= ChargeStatusDetail(len(_ChargeStatusDetail_index)-1) {
		return fmt.Sprintf("ChargeStatusDetail(%d)", i)
	}
	return _ChargeStatusDetail_name[_ChargeStatusDetail_index[i]:_ChargeStatusDetail_index[i+1]]
}
