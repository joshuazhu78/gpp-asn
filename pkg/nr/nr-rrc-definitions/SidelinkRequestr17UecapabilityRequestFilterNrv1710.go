package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s SidelinkRequestr17UecapabilityRequestFilterNrv1710) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *SidelinkRequestr17UecapabilityRequestFilterNrv1710) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseSidelinkRequestr17UecapabilityRequestFilterNrv1710(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseSidelinkRequestr17UecapabilityRequestFilterNrv1710(s string) (SidelinkRequestr17UecapabilityRequestFilterNrv1710, error) {
	value, ok := SidelinkRequestr17UecapabilityRequestFilterNrv1710_value[s]
	if !ok {
		return SidelinkRequestr17UecapabilityRequestFilterNrv1710(0), fmt.Errorf("%q is not a valid %s", s, "SidelinkRequestr17UecapabilityRequestFilterNrv1710")
	}
	return SidelinkRequestr17UecapabilityRequestFilterNrv1710(value), nil
}
