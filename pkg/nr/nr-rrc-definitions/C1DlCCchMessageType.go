package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var C1DlCcchMessageTypeknownImplementations = []isC1DlCCchMessageType_C1DlCcchMessageType{
	&C1DlCCchMessageType_RrcReject{},
	&C1DlCCchMessageType_RrcSetup{},
	&C1DlCCchMessageType_Spare2{},
	&C1DlCCchMessageType_Spare1{},
}

func (c *C1DlCCchMessageType) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range C1DlCcchMessageTypeknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.C1DlCcchMessageType = target.Elem().Interface().(isC1DlCCchMessageType_C1DlCcchMessageType)
			return nil
		}
	}
	return nil
}

func (c C1DlCCchMessageType) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.C1DlCcchMessageType).String(),
		Value: c.C1DlCcchMessageType,
	})
}

var _ json.Unmarshaler = (*C1DlCCchMessageType)(nil)
