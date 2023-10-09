package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var N1N2MoreThanTwoknownImplementations = []isN1N2MoreThanTwo_N1N2MoreThanTwo{
	&N1N2MoreThanTwo_TwoOneTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_TwoTwoTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_FourOneTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_ThreeTwoTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_SixOneTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_FourTwoTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_EightOneTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_FourThreeTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_SixTwoTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_TwelveOneTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_FourFourTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_EightTwoTypeISinglePanelRestriction{},
	&N1N2MoreThanTwo_SixteenOneTypeISinglePanelRestriction{},
}

func (c *N1N2MoreThanTwo) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range N1N2MoreThanTwoknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.N1N2MoreThanTwo = target.Elem().Interface().(isN1N2MoreThanTwo_N1N2MoreThanTwo)
			return nil
		}
	}
	return nil
}

func (c N1N2MoreThanTwo) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.N1N2MoreThanTwo).String(),
		Value: c.N1N2MoreThanTwo,
	})
}

var _ json.Unmarshaler = (*N1N2MoreThanTwo)(nil)
