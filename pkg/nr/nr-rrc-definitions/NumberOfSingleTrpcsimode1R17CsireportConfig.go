package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s NumberOfSingleTrpcsimode1R17CsireportConfig) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *NumberOfSingleTrpcsimode1R17CsireportConfig) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseNumberOfSingleTrpcsimode1R17CsireportConfig(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseNumberOfSingleTrpcsimode1R17CsireportConfig(s string) (NumberOfSingleTrpcsimode1R17CsireportConfig, error) {
	value, ok := NumberOfSingleTrpcsimode1R17CsireportConfig_value[s]
	if !ok {
		return NumberOfSingleTrpcsimode1R17CsireportConfig(0), fmt.Errorf("%q is not a valid %s", s, "NumberOfSingleTrpcsimode1R17CsireportConfig")
	}
	return NumberOfSingleTrpcsimode1R17CsireportConfig(value), nil
}
