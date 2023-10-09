package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var PortIndexFor8RanksknownImplementations = []isPortIndexFor8Ranks_PortIndexFor8Ranks{
	&PortIndexFor8Ranks_PortIndex8{},
	&PortIndexFor8Ranks_PortIndex4{},
	&PortIndexFor8Ranks_PortIndex2{},
	&PortIndexFor8Ranks_PortIndex1{},
}

func (c *PortIndexFor8Ranks) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range PortIndexFor8RanksknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.PortIndexFor8Ranks = target.Elem().Interface().(isPortIndexFor8Ranks_PortIndexFor8Ranks)
			return nil
		}
	}
	return nil
}

func (c PortIndexFor8Ranks) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.PortIndexFor8Ranks).String(),
		Value: c.PortIndexFor8Ranks,
	})
}

var _ json.Unmarshaler = (*PortIndexFor8Ranks)(nil)
