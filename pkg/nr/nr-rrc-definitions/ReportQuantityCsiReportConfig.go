package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var ReportQuantityCsiReportConfigknownImplementations = []isReportQuantityCsiReportConfig_ReportQuantityCsiReportConfig{
	&ReportQuantityCsiReportConfig_None{},
	&ReportQuantityCsiReportConfig_CriRiPmiCqi{},
	&ReportQuantityCsiReportConfig_CriRiI1{},
	&ReportQuantityCsiReportConfig_CriRiI1Cqi{},
	&ReportQuantityCsiReportConfig_CriRiCqi{},
	&ReportQuantityCsiReportConfig_CriRsrp{},
	&ReportQuantityCsiReportConfig_SsbIndexRsrp{},
	&ReportQuantityCsiReportConfig_CriRiLiPmiCqi{},
}

func (c *ReportQuantityCsiReportConfig) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range ReportQuantityCsiReportConfigknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.ReportQuantityCsiReportConfig = target.Elem().Interface().(isReportQuantityCsiReportConfig_ReportQuantityCsiReportConfig)
			return nil
		}
	}
	return nil
}

func (c ReportQuantityCsiReportConfig) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.ReportQuantityCsiReportConfig).String(),
		Value: c.ReportQuantityCsiReportConfig,
	})
}

var _ json.Unmarshaler = (*ReportQuantityCsiReportConfig)(nil)
