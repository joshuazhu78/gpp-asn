package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var CsiResourcePeriodicityAndOffsetknownImplementations = []isCsiResourcePeriodicityAndOffset_CsiResourcePeriodicityAndOffset{
	&CsiResourcePeriodicityAndOffset_Slots4{},
	&CsiResourcePeriodicityAndOffset_Slots5{},
	&CsiResourcePeriodicityAndOffset_Slots8{},
	&CsiResourcePeriodicityAndOffset_Slots10{},
	&CsiResourcePeriodicityAndOffset_Slots16{},
	&CsiResourcePeriodicityAndOffset_Slots20{},
	&CsiResourcePeriodicityAndOffset_Slots32{},
	&CsiResourcePeriodicityAndOffset_Slots40{},
	&CsiResourcePeriodicityAndOffset_Slots64{},
	&CsiResourcePeriodicityAndOffset_Slots80{},
	&CsiResourcePeriodicityAndOffset_Slots160{},
	&CsiResourcePeriodicityAndOffset_Slots320{},
	&CsiResourcePeriodicityAndOffset_Slots640{},
}

func (c *CsiResourcePeriodicityAndOffset) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range CsiResourcePeriodicityAndOffsetknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.CsiResourcePeriodicityAndOffset = target.Elem().Interface().(isCsiResourcePeriodicityAndOffset_CsiResourcePeriodicityAndOffset)
			return nil
		}
	}
	return nil
}

func (c CsiResourcePeriodicityAndOffset) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.CsiResourcePeriodicityAndOffset).String(),
		Value: c.CsiResourcePeriodicityAndOffset,
	})
}

var _ json.Unmarshaler = (*CsiResourcePeriodicityAndOffset)(nil)
