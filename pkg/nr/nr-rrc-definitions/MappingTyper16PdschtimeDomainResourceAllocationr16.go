package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s MappingTyper16PdschtimeDomainResourceAllocationr16) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *MappingTyper16PdschtimeDomainResourceAllocationr16) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseMappingTyper16PdschtimeDomainResourceAllocationr16(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseMappingTyper16PdschtimeDomainResourceAllocationr16(s string) (MappingTyper16PdschtimeDomainResourceAllocationr16, error) {
	value, ok := MappingTyper16PdschtimeDomainResourceAllocationr16_value[s]
	if !ok {
		return MappingTyper16PdschtimeDomainResourceAllocationr16(0), fmt.Errorf("%q is not a valid %s", s, "MappingTyper16PdschtimeDomainResourceAllocationr16")
	}
	return MappingTyper16PdschtimeDomainResourceAllocationr16(value), nil
}
