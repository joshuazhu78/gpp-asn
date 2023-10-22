package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s PhaseAlphabetSizetype2) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *PhaseAlphabetSizetype2) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParsePhaseAlphabetSizetype2(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParsePhaseAlphabetSizetype2(s string) (PhaseAlphabetSizetype2, error) {
	value, ok := PhaseAlphabetSizetype2_value[s]
	if !ok {
		return PhaseAlphabetSizetype2(0), fmt.Errorf("%q is not a valid %s", s, "PhaseAlphabetSizetype2")
	}
	return PhaseAlphabetSizetype2(value), nil
}
