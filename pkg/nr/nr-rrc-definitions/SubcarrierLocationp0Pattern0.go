package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s SubcarrierLocationp0Pattern0) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *SubcarrierLocationp0Pattern0) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseSubcarrierLocationp0Pattern0(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseSubcarrierLocationp0Pattern0(s string) (SubcarrierLocationp0Pattern0, error) {
	value, ok := SubcarrierLocationp0Pattern0_value[s]
	if !ok {
		return SubcarrierLocationp0Pattern0(0), fmt.Errorf("%q is not a valid %s", s, "SubcarrierLocationp0Pattern0")
	}
	return SubcarrierLocationp0Pattern0(value), nil
}