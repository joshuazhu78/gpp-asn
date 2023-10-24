package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseMaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseMaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16(s string) (MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16, error) {
	value, ok := MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16_value[s]
	if !ok {
		return MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16(0), fmt.Errorf("%q is not a valid %s", s, "MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16")
	}
	return MaxNumberPerCoresetpoolr16MaxNumberActivatedTcistatesr16(value), nil
}
