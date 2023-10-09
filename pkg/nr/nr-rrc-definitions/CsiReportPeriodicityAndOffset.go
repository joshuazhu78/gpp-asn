package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var CsiReportPeriodicityAndOffsetknownImplementations = []isCsiReportPeriodicityAndOffset_CsiReportPeriodicityAndOffset{
	&CsiReportPeriodicityAndOffset_Slots4{},
	&CsiReportPeriodicityAndOffset_Slots5{},
	&CsiReportPeriodicityAndOffset_Slots8{},
	&CsiReportPeriodicityAndOffset_Slots10{},
	&CsiReportPeriodicityAndOffset_Slots16{},
	&CsiReportPeriodicityAndOffset_Slots20{},
	&CsiReportPeriodicityAndOffset_Slots40{},
	&CsiReportPeriodicityAndOffset_Slots80{},
	&CsiReportPeriodicityAndOffset_Slots160{},
	&CsiReportPeriodicityAndOffset_Slots320{},
}

func (c *CsiReportPeriodicityAndOffset) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range CsiReportPeriodicityAndOffsetknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.CsiReportPeriodicityAndOffset = target.Elem().Interface().(isCsiReportPeriodicityAndOffset_CsiReportPeriodicityAndOffset)
			return nil
		}
	}
	return nil
}

func (c CsiReportPeriodicityAndOffset) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.CsiReportPeriodicityAndOffset).String(),
		Value: c.CsiReportPeriodicityAndOffset,
	})
}

var _ json.Unmarshaler = (*CsiReportPeriodicityAndOffset)(nil)
