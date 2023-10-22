package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s DlUltransmissionPeriodicityv1530Tdduldlpattern) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *DlUltransmissionPeriodicityv1530Tdduldlpattern) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseDlUltransmissionPeriodicityv1530Tdduldlpattern(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseDlUltransmissionPeriodicityv1530Tdduldlpattern(s string) (DlUltransmissionPeriodicityv1530Tdduldlpattern, error) {
	value, ok := DlUltransmissionPeriodicityv1530Tdduldlpattern_value[s]
	if !ok {
		return DlUltransmissionPeriodicityv1530Tdduldlpattern(0), fmt.Errorf("%q is not a valid %s", s, "DlUltransmissionPeriodicityv1530Tdduldlpattern")
	}
	return DlUltransmissionPeriodicityv1530Tdduldlpattern(value), nil
}
