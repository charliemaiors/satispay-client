package client

import (
	"encoding/json"
	"fmt"
)

var (
	_ChargeStatusDetailNameToValue = map[string]ChargeStatusDetail{
		"DECLINED_BY_PAYER":              DeclinedByPayer,
		"DECLINED_BY_PAYER_NOT_REQUIRED": DeclinedByPayerNotRequired,
		"CANCEL_BY_NEW_CHARGE":           CancelByNewCharge,
		"INTERNAL_FAILURE":               InternalFailure,
		"EXPIRED":                        Expired,
	}

	_ChargeStatusDetailValueToName = map[ChargeStatusDetail]string{
		DeclinedByPayer:            "DECLINED_BY_PAYER",
		DeclinedByPayerNotRequired: "DECLINED_BY_PAYER_NOT_REQUIRED",
		CancelByNewCharge:          "CANCEL_BY_NEW_CHARGE",
		InternalFailure:            "INTERNAL_FAILURE",
		Expired:                    "EXPIRED",
	}
)

func init() {
	var v ChargeStatusDetail
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ChargeStatusDetailNameToValue = map[string]ChargeStatusDetail{
			interface{}(DeclinedByPayer).(fmt.Stringer).String():            DeclinedByPayer,
			interface{}(DeclinedByPayerNotRequired).(fmt.Stringer).String(): DeclinedByPayerNotRequired,
			interface{}(CancelByNewCharge).(fmt.Stringer).String():          CancelByNewCharge,
			interface{}(InternalFailure).(fmt.Stringer).String():            InternalFailure,
			interface{}(Expired).(fmt.Stringer).String():                    Expired,
		}
	}
}

// MarshalJSON is generated so ChargeStatusDetail satisfies json.Marshaler.
func (r ChargeStatusDetail) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ChargeStatusDetailValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ChargeStatusDetail: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ChargeStatusDetail satisfies json.Unmarshaler.
func (r *ChargeStatusDetail) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ChargeStatusDetail should be a string, got %s", data)
	}
	v, ok := _ChargeStatusDetailNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ChargeStatusDetail %q", s)
	}
	*r = v
	return nil
}
