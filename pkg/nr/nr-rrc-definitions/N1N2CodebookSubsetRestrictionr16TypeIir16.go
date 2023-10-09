package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var N1N2CodebookSubsetRestrictionR16TypeIiR16knownImplementations = []isN1N2CodebookSubsetRestrictionr16TypeIir16_N1N2CodebookSubsetRestrictionR16TypeIiR16{
	&N1N2CodebookSubsetRestrictionr16TypeIir16_TwoOne{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_TwoTwo{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_FourOne{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_ThreeTwo{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_SixOne{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_FourTwo{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_EightOne{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_FourThree{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_SixTwo{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_TwelveOne{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_FourFour{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_EightTwo{},
	&N1N2CodebookSubsetRestrictionr16TypeIir16_SixteenOne{},
}

func (c *N1N2CodebookSubsetRestrictionr16TypeIir16) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range N1N2CodebookSubsetRestrictionR16TypeIiR16knownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.N1N2CodebookSubsetRestrictionR16TypeIiR16 = target.Elem().Interface().(isN1N2CodebookSubsetRestrictionr16TypeIir16_N1N2CodebookSubsetRestrictionR16TypeIiR16)
			return nil
		}
	}
	return nil
}

func (c N1N2CodebookSubsetRestrictionr16TypeIir16) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.N1N2CodebookSubsetRestrictionR16TypeIiR16).String(),
		Value: c.N1N2CodebookSubsetRestrictionR16TypeIiR16,
	})
}

var _ json.Unmarshaler = (*N1N2CodebookSubsetRestrictionr16TypeIir16)(nil)
