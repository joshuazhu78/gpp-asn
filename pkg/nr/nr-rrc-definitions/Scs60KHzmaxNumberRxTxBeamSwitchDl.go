package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s Scs60KHzmaxNumberRxTxBeamSwitchDl) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *Scs60KHzmaxNumberRxTxBeamSwitchDl) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseScs60KHzmaxNumberRxTxBeamSwitchDl(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseScs60KHzmaxNumberRxTxBeamSwitchDl(s string) (Scs60KHzmaxNumberRxTxBeamSwitchDl, error) {
	value, ok := Scs60KHzmaxNumberRxTxBeamSwitchDl_value[s]
	if !ok {
		return Scs60KHzmaxNumberRxTxBeamSwitchDl(0), fmt.Errorf("%q is not a valid %s", s, "Scs60KHzmaxNumberRxTxBeamSwitchDl")
	}
	return Scs60KHzmaxNumberRxTxBeamSwitchDl(value), nil
}
