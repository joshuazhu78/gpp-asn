package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s PdschBundleSizeForCsicriRii1Cqi) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *PdschBundleSizeForCsicriRii1Cqi) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParsePdschBundleSizeForCsicriRii1Cqi(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParsePdschBundleSizeForCsicriRii1Cqi(s string) (PdschBundleSizeForCsicriRii1Cqi, error) {
	value, ok := PdschBundleSizeForCsicriRii1Cqi_value[s]
	if !ok {
		return PdschBundleSizeForCsicriRii1Cqi(0), fmt.Errorf("%q is not a valid %s", s, "PdschBundleSizeForCsicriRii1Cqi")
	}
	return PdschBundleSizeForCsicriRii1Cqi(value), nil
}
