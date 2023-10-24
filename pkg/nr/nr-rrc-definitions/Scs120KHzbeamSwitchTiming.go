package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s Scs120KHzbeamSwitchTiming) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *Scs120KHzbeamSwitchTiming) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseScs120KHzbeamSwitchTiming(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseScs120KHzbeamSwitchTiming(s string) (Scs120KHzbeamSwitchTiming, error) {
	value, ok := Scs120KHzbeamSwitchTiming_value[s]
	if !ok {
		return Scs120KHzbeamSwitchTiming(0), fmt.Errorf("%q is not a valid %s", s, "Scs120KHzbeamSwitchTiming")
	}
	return Scs120KHzbeamSwitchTiming(value), nil
}
