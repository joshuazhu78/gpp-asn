package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var SubTypetype2knownImplementations = []isSubTypetype2_SubTypetype2{
	&SubTypetype2_TypeIi{},
	&SubTypetype2_TypeIiPortSelection{},
}

func (c *SubTypetype2) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range SubTypetype2knownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.SubTypetype2 = target.Elem().Interface().(isSubTypetype2_SubTypetype2)
			return nil
		}
	}
	return nil
}

func (c SubTypetype2) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.SubTypetype2).String(),
		Value: c.SubTypetype2,
	})
}

var _ json.Unmarshaler = (*SubTypetype2)(nil)
