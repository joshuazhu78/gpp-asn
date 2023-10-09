package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var CsiReportingBandreportFreqConfigurationknownImplementations = []isCsiReportingBandreportFreqConfiguration_CsiReportingBandreportFreqConfiguration{
	&CsiReportingBandreportFreqConfiguration_Subbands3{},
	&CsiReportingBandreportFreqConfiguration_Subbands4{},
	&CsiReportingBandreportFreqConfiguration_Subbands5{},
	&CsiReportingBandreportFreqConfiguration_Subbands6{},
	&CsiReportingBandreportFreqConfiguration_Subbands7{},
	&CsiReportingBandreportFreqConfiguration_Subbands8{},
	&CsiReportingBandreportFreqConfiguration_Subbands9{},
	&CsiReportingBandreportFreqConfiguration_Subbands10{},
	&CsiReportingBandreportFreqConfiguration_Subbands11{},
	&CsiReportingBandreportFreqConfiguration_Subbands12{},
	&CsiReportingBandreportFreqConfiguration_Subbands13{},
	&CsiReportingBandreportFreqConfiguration_Subbands14{},
	&CsiReportingBandreportFreqConfiguration_Subbands15{},
	&CsiReportingBandreportFreqConfiguration_Subbands16{},
	&CsiReportingBandreportFreqConfiguration_Subbands17{},
	&CsiReportingBandreportFreqConfiguration_Subbands18{},
	&CsiReportingBandreportFreqConfiguration_Subbands19V1530{},
}

func (c *CsiReportingBandreportFreqConfiguration) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range CsiReportingBandreportFreqConfigurationknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.CsiReportingBandreportFreqConfiguration = target.Elem().Interface().(isCsiReportingBandreportFreqConfiguration_CsiReportingBandreportFreqConfiguration)
			return nil
		}
	}
	return nil
}

func (c CsiReportingBandreportFreqConfiguration) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.CsiReportingBandreportFreqConfiguration).String(),
		Value: c.CsiReportingBandreportFreqConfiguration,
	})
}

var _ json.Unmarshaler = (*CsiReportingBandreportFreqConfiguration)(nil)
