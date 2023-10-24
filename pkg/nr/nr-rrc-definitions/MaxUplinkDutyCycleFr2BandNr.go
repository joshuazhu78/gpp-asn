package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s MaxUplinkDutyCycleFr2BandNr) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *MaxUplinkDutyCycleFr2BandNr) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseMaxUplinkDutyCycleFr2BandNr(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseMaxUplinkDutyCycleFr2BandNr(s string) (MaxUplinkDutyCycleFr2BandNr, error) {
	value, ok := MaxUplinkDutyCycleFr2BandNr_value[s]
	if !ok {
		return MaxUplinkDutyCycleFr2BandNr(0), fmt.Errorf("%q is not a valid %s", s, "MaxUplinkDutyCycleFr2BandNr")
	}
	return MaxUplinkDutyCycleFr2BandNr(value), nil
}
