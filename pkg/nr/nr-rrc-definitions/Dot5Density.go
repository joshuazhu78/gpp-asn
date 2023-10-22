package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s Dot5Density) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *Dot5Density) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseDot5Density(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseDot5Density(s string) (Dot5Density, error) {
	value, ok := Dot5Density_value[s]
	if !ok {
		return Dot5Density(0), fmt.Errorf("%q is not a valid %s", s, "Dot5Density")
	}
	return Dot5Density(value), nil
}
