package client

import (
	"encoding/json"
	"fmt"
)

var (
	_RefundReasonNameToValue = map[string]RefundReason{
		"DUPLICATE":             Duplicate,
		"FRAUDOLENT":            Fraudolent,
		"REQUESTED_BY_CUSTOMER": RequestedByCustomer,
	}

	_RefundReasonValueToName = map[RefundReason]string{
		Duplicate:           "DUPLICATE",
		Fraudolent:          "FRAUDOLENT",
		RequestedByCustomer: "REQUESTED_BY_CUSTOMER",
	}
)

func init() {
	var v RefundReason
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_RefundReasonNameToValue = map[string]RefundReason{
			interface{}(Duplicate).(fmt.Stringer).String():           Duplicate,
			interface{}(Fraudolent).(fmt.Stringer).String():          Fraudolent,
			interface{}(RequestedByCustomer).(fmt.Stringer).String(): RequestedByCustomer,
		}
	}
}

// MarshalJSON is generated so RefundReason satisfies json.Marshaler.
func (r RefundReason) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _RefundReasonValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid RefundReason: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so RefundReason satisfies json.Unmarshaler.
func (r *RefundReason) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RefundReason should be a string, got %s", data)
	}
	v, ok := _RefundReasonNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid RefundReason %q", s)
	}
	*r = v
	return nil
}
