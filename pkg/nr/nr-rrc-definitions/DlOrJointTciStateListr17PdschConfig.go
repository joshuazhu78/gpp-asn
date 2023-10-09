package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var DlOrJointTciStateListR17PdschConfigknownImplementations = []isDlOrJointTciStateListr17PdschConfig_DlOrJointTciStateListR17PdschConfig{
	&DlOrJointTciStateListr17PdschConfig_Explicitlist{},
	&DlOrJointTciStateListr17PdschConfig_UnifiedTciStateRefR17{},
}

func (c *DlOrJointTciStateListr17PdschConfig) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range DlOrJointTciStateListR17PdschConfigknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.DlOrJointTciStateListR17PdschConfig = target.Elem().Interface().(isDlOrJointTciStateListr17PdschConfig_DlOrJointTciStateListR17PdschConfig)
			return nil
		}
	}
	return nil
}

func (c DlOrJointTciStateListr17PdschConfig) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.DlOrJointTciStateListR17PdschConfig).String(),
		Value: c.DlOrJointTciStateListR17PdschConfig,
	})
}

var _ json.Unmarshaler = (*DlOrJointTciStateListr17PdschConfig)(nil)
