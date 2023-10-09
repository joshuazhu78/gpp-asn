package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var C1DlDcchMessageTypeknownImplementations = []isC1DlDCchMessageType_C1DlDcchMessageType{
	&C1DlDCchMessageType_RrcReconfiguration{},
	&C1DlDCchMessageType_RrcResume{},
	&C1DlDCchMessageType_RrcRelease{},
	&C1DlDCchMessageType_RrcReestablishment{},
	&C1DlDCchMessageType_SecurityModeCommand{},
	&C1DlDCchMessageType_DlInformationTransfer{},
	&C1DlDCchMessageType_UeCapabilityEnquiry{},
	&C1DlDCchMessageType_CounterCheck{},
	&C1DlDCchMessageType_MobilityFromNrcommand{},
	&C1DlDCchMessageType_DlDedicatedMessageSegmentR16{},
	&C1DlDCchMessageType_UeInformationRequestR16{},
	&C1DlDCchMessageType_DlInformationTransferMrdcR16{},
	&C1DlDCchMessageType_LoggedMeasurementConfigurationR16{},
	&C1DlDCchMessageType_Spare3{},
	&C1DlDCchMessageType_Spare2{},
	&C1DlDCchMessageType_Spare1{},
}

func (c *C1DlDCchMessageType) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range C1DlDcchMessageTypeknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.C1DlDcchMessageType = target.Elem().Interface().(isC1DlDCchMessageType_C1DlDcchMessageType)
			return nil
		}
	}
	return nil
}

func (c C1DlDCchMessageType) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.C1DlDcchMessageType).String(),
		Value: c.C1DlDcchMessageType,
	})
}

var _ json.Unmarshaler = (*C1DlDCchMessageType)(nil)
