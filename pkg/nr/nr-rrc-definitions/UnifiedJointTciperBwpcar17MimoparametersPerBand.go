package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s UnifiedJointTciperBwpcar17MimoparametersPerBand) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *UnifiedJointTciperBwpcar17MimoparametersPerBand) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseUnifiedJointTciperBwpcar17MimoparametersPerBand(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseUnifiedJointTciperBwpcar17MimoparametersPerBand(s string) (UnifiedJointTciperBwpcar17MimoparametersPerBand, error) {
	value, ok := UnifiedJointTciperBwpcar17MimoparametersPerBand_value[s]
	if !ok {
		return UnifiedJointTciperBwpcar17MimoparametersPerBand(0), fmt.Errorf("%q is not a valid %s", s, "UnifiedJointTciperBwpcar17MimoparametersPerBand")
	}
	return UnifiedJointTciperBwpcar17MimoparametersPerBand(value), nil
}
