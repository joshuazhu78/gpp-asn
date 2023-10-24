package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s UnifiedJointTcilegacySrsr17MimoparametersPerBand) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *UnifiedJointTcilegacySrsr17MimoparametersPerBand) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseUnifiedJointTcilegacySrsr17MimoparametersPerBand(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseUnifiedJointTcilegacySrsr17MimoparametersPerBand(s string) (UnifiedJointTcilegacySrsr17MimoparametersPerBand, error) {
	value, ok := UnifiedJointTcilegacySrsr17MimoparametersPerBand_value[s]
	if !ok {
		return UnifiedJointTcilegacySrsr17MimoparametersPerBand(0), fmt.Errorf("%q is not a valid %s", s, "UnifiedJointTcilegacySrsr17MimoparametersPerBand")
	}
	return UnifiedJointTcilegacySrsr17MimoparametersPerBand(value), nil
}
