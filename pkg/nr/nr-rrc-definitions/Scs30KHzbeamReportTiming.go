package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s Scs30KHzbeamReportTiming) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *Scs30KHzbeamReportTiming) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseScs30KHzbeamReportTiming(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseScs30KHzbeamReportTiming(s string) (Scs30KHzbeamReportTiming, error) {
	value, ok := Scs30KHzbeamReportTiming_value[s]
	if !ok {
		return Scs30KHzbeamReportTiming(0), fmt.Errorf("%q is not a valid %s", s, "Scs30KHzbeamReportTiming")
	}
	return Scs30KHzbeamReportTiming(value), nil
}
