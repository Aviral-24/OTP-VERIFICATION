package data

import (
	"encoding/json"
	"strconv"
)

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}

// Custom unmarshaler to handle phoneNumber as both string and number
func (o *OTPData) UnmarshalJSON(data []byte) error {
	type Alias OTPData
	aux := &struct {
		PhoneNumber interface{} `json:"phoneNumber"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.PhoneNumber.(type) {
	case string:
		o.PhoneNumber = v
	case float64:
		o.PhoneNumber = strconv.FormatFloat(v, 'f', 0, 64)
	}
	return nil
}

func (v *VerifyData) UnmarshalJSON(data []byte) error {
	type Alias VerifyData
	aux := &struct {
		PhoneNumber interface{} `json:"phoneNumber"`
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch p := aux.PhoneNumber.(type) {
	case string:
		v.PhoneNumber = p
	case float64:
		v.PhoneNumber = strconv.FormatFloat(p, 'f', 0, 64)
	}
	return nil
}
