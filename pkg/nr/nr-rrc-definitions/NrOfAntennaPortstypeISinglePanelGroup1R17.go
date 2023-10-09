package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var NrOfAntennaPortstypeISinglePanelGroup1R17knownImplementations = []isNrOfAntennaPortstypeISinglePanelGroup1R17_NrOfAntennaPortstypeISinglePanelGroup1R17{
	&NrOfAntennaPortstypeISinglePanelGroup1R17_Two{},
	&NrOfAntennaPortstypeISinglePanelGroup1R17_MoreThanTwo{},
}

func (c *NrOfAntennaPortstypeISinglePanelGroup1R17) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range NrOfAntennaPortstypeISinglePanelGroup1R17knownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.NrOfAntennaPortstypeISinglePanelGroup1R17 = target.Elem().Interface().(isNrOfAntennaPortstypeISinglePanelGroup1R17_NrOfAntennaPortstypeISinglePanelGroup1R17)
			return nil
		}
	}
	return nil
}

func (c NrOfAntennaPortstypeISinglePanelGroup1R17) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.NrOfAntennaPortstypeISinglePanelGroup1R17).String(),
		Value: c.NrOfAntennaPortstypeISinglePanelGroup1R17,
	})
}

var _ json.Unmarshaler = (*NrOfAntennaPortstypeISinglePanelGroup1R17)(nil)
