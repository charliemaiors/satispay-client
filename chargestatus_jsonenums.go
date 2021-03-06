// generated by jsonenums -type=ChargeStatus; DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"
)

var (
	_ChargeStatusNameToValue = map[string]ChargeStatus{
		"REQUIRED": Required,
		"SUCCESS":  Success,
		"FAILURE":  Failure,
	}

	_ChargeStatusValueToName = map[ChargeStatus]string{
		Required: "REQUIRED",
		Success:  "SUCCESS",
		Failure:  "FAILURE",
	}
)

func init() {
	var v ChargeStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ChargeStatusNameToValue = map[string]ChargeStatus{
			interface{}(Required).(fmt.Stringer).String(): Required,
			interface{}(Success).(fmt.Stringer).String():  Success,
			interface{}(Failure).(fmt.Stringer).String():  Failure,
		}
	}
}

// MarshalJSON is generated so ChargeStatus satisfies json.Marshaler.
func (r ChargeStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ChargeStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ChargeStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ChargeStatus satisfies json.Unmarshaler.
func (r *ChargeStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ChargeStatus should be a string, got %s", data)
	}
	v, ok := _ChargeStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ChargeStatus %q", s)
	}
	*r = v
	return nil
}
