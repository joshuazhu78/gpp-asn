package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s MaxConfiguredDltcir17UnifiedSeparateTcir17) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *MaxConfiguredDltcir17UnifiedSeparateTcir17) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseMaxConfiguredDltcir17UnifiedSeparateTcir17(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseMaxConfiguredDltcir17UnifiedSeparateTcir17(s string) (MaxConfiguredDltcir17UnifiedSeparateTcir17, error) {
	value, ok := MaxConfiguredDltcir17UnifiedSeparateTcir17_value[s]
	if !ok {
		return MaxConfiguredDltcir17UnifiedSeparateTcir17(0), fmt.Errorf("%q is not a valid %s", s, "MaxConfiguredDltcir17UnifiedSeparateTcir17")
	}
	return MaxConfiguredDltcir17UnifiedSeparateTcir17(value), nil
}
