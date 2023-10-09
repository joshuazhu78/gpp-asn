package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var ReportQuantityR16CsiReportConfigknownImplementations = []isReportQuantityr16CsiReportConfig_ReportQuantityR16CsiReportConfig{
	&ReportQuantityr16CsiReportConfig_CriSinrR16{},
	&ReportQuantityr16CsiReportConfig_SsbIndexSinrR16{},
}

func (c *ReportQuantityr16CsiReportConfig) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range ReportQuantityR16CsiReportConfigknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.ReportQuantityR16CsiReportConfig = target.Elem().Interface().(isReportQuantityr16CsiReportConfig_ReportQuantityR16CsiReportConfig)
			return nil
		}
	}
	return nil
}

func (c ReportQuantityr16CsiReportConfig) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.ReportQuantityR16CsiReportConfig).String(),
		Value: c.ReportQuantityR16CsiReportConfig,
	})
}

var _ json.Unmarshaler = (*ReportQuantityr16CsiReportConfig)(nil)
