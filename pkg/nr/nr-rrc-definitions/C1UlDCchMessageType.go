package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var C1UlDcchMessageTypeknownImplementations = []isC1UlDCchMessageType_C1UlDcchMessageType{
	&C1UlDCchMessageType_MeasurementReport{},
	&C1UlDCchMessageType_RrcReconfigurationComplete{},
	&C1UlDCchMessageType_RrcSetupComplete{},
	&C1UlDCchMessageType_RrcReestablishmentComplete{},
	&C1UlDCchMessageType_RrcResumeComplete{},
	&C1UlDCchMessageType_SecurityModeComplete{},
	&C1UlDCchMessageType_SecurityModeFailure{},
	&C1UlDCchMessageType_UlInformationTransfer{},
	&C1UlDCchMessageType_LocationMeasurementIndication{},
	&C1UlDCchMessageType_UeCapabilityInformation{},
	&C1UlDCchMessageType_CounterCheckResponse{},
	&C1UlDCchMessageType_UeAssistanceInformation{},
	&C1UlDCchMessageType_FailureInformation{},
	&C1UlDCchMessageType_UlInformationTransferMrdc{},
	&C1UlDCchMessageType_ScgFailureInformation{},
	&C1UlDCchMessageType_ScgFailureInformationEutra{},
}

func (c *C1UlDCchMessageType) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range C1UlDcchMessageTypeknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.C1UlDcchMessageType = target.Elem().Interface().(isC1UlDCchMessageType_C1UlDcchMessageType)
			return nil
		}
	}
	return nil
}

func (c C1UlDCchMessageType) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.C1UlDcchMessageType).String(),
		Value: c.C1UlDcchMessageType,
	})
}

var _ json.Unmarshaler = (*C1UlDCchMessageType)(nil)
