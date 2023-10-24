package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s BeamCorrespondenceSsbbasedr16MimoparametersPerBand) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *BeamCorrespondenceSsbbasedr16MimoparametersPerBand) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseBeamCorrespondenceSsbbasedr16MimoparametersPerBand(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseBeamCorrespondenceSsbbasedr16MimoparametersPerBand(s string) (BeamCorrespondenceSsbbasedr16MimoparametersPerBand, error) {
	value, ok := BeamCorrespondenceSsbbasedr16MimoparametersPerBand_value[s]
	if !ok {
		return BeamCorrespondenceSsbbasedr16MimoparametersPerBand(0), fmt.Errorf("%q is not a valid %s", s, "BeamCorrespondenceSsbbasedr16MimoparametersPerBand")
	}
	return BeamCorrespondenceSsbbasedr16MimoparametersPerBand(value), nil
}
