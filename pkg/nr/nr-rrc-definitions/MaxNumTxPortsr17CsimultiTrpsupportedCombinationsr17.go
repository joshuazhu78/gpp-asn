package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseMaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseMaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17(s string) (MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17, error) {
	value, ok := MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17_value[s]
	if !ok {
		return MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17(0), fmt.Errorf("%q is not a valid %s", s, "MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17")
	}
	return MaxNumTxPortsr17CsimultiTrpsupportedCombinationsr17(value), nil
}
