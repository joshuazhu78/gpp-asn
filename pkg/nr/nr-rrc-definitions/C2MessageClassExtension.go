package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var C2MessageClassExtensionknownImplementations = []isC2MessageClassExtension_C2MessageClassExtension{
	&C2MessageClassExtension_UlDedicatedMessageSegmentR16{},
	&C2MessageClassExtension_DedicatedSibrequestR16{},
	&C2MessageClassExtension_McgFailureInformationR16{},
	&C2MessageClassExtension_UeInformationResponseR16{},
	&C2MessageClassExtension_SidelinkUeinformationNrR16{},
	&C2MessageClassExtension_UlInformationTransferIratR16{},
	&C2MessageClassExtension_IabOtherInformationR16{},
	&C2MessageClassExtension_MbsInterestIndicationR17{},
	&C2MessageClassExtension_UePositioningAssistanceInfoR17{},
	&C2MessageClassExtension_MeasurementReportAppLayerR17{},
	&C2MessageClassExtension_Spare6{},
	&C2MessageClassExtension_Spare5{},
	&C2MessageClassExtension_Spare4{},
	&C2MessageClassExtension_Spare3{},
	&C2MessageClassExtension_Spare2{},
	&C2MessageClassExtension_Spare1{},
}

func (c *C2MessageClassExtension) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range C2MessageClassExtensionknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.C2MessageClassExtension = target.Elem().Interface().(isC2MessageClassExtension_C2MessageClassExtension)
			return nil
		}
	}
	return nil
}

func (c C2MessageClassExtension) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.C2MessageClassExtension).String(),
		Value: c.C2MessageClassExtension,
	})
}

var _ json.Unmarshaler = (*C2MessageClassExtension)(nil)
