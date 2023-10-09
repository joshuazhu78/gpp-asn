package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var NgN1N2TypeIMultiPanelknownImplementations = []isNgn1N2TypeIMultiPanel_NgN1N2TypeIMultiPanel{
	&Ngn1N2TypeIMultiPanel_TwoTwoOneTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_TwoFourOneTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_FourTwoOneTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_TwoTwoTwoTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_TwoEightOneTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_FourFourOneTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_TwoFourTwoTypeIMultiPanelRestriction{},
	&Ngn1N2TypeIMultiPanel_FourTwoTwoTypeIMultiPanelRestriction{},
}

func (c *Ngn1N2TypeIMultiPanel) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range NgN1N2TypeIMultiPanelknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.NgN1N2TypeIMultiPanel = target.Elem().Interface().(isNgn1N2TypeIMultiPanel_NgN1N2TypeIMultiPanel)
			return nil
		}
	}
	return nil
}

func (c Ngn1N2TypeIMultiPanel) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.NgN1N2TypeIMultiPanel).String(),
		Value: c.NgN1N2TypeIMultiPanel,
	})
}

var _ json.Unmarshaler = (*Ngn1N2TypeIMultiPanel)(nil)
