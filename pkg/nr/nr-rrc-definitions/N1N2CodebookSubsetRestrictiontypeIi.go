package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var N1N2CodebookSubsetRestrictiontypeIiknownImplementations = []isN1N2CodebookSubsetRestrictiontypeIi_N1N2CodebookSubsetRestrictiontypeIi{
	&N1N2CodebookSubsetRestrictiontypeIi_TwoOne{},
	&N1N2CodebookSubsetRestrictiontypeIi_TwoTwo{},
	&N1N2CodebookSubsetRestrictiontypeIi_FourOne{},
	&N1N2CodebookSubsetRestrictiontypeIi_ThreeTwo{},
	&N1N2CodebookSubsetRestrictiontypeIi_SixOne{},
	&N1N2CodebookSubsetRestrictiontypeIi_FourTwo{},
	&N1N2CodebookSubsetRestrictiontypeIi_EightOne{},
	&N1N2CodebookSubsetRestrictiontypeIi_FourThree{},
	&N1N2CodebookSubsetRestrictiontypeIi_SixTwo{},
	&N1N2CodebookSubsetRestrictiontypeIi_TwelveOne{},
	&N1N2CodebookSubsetRestrictiontypeIi_FourFour{},
	&N1N2CodebookSubsetRestrictiontypeIi_EightTwo{},
	&N1N2CodebookSubsetRestrictiontypeIi_SixteenOne{},
}

func (c *N1N2CodebookSubsetRestrictiontypeIi) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range N1N2CodebookSubsetRestrictiontypeIiknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.N1N2CodebookSubsetRestrictiontypeIi = target.Elem().Interface().(isN1N2CodebookSubsetRestrictiontypeIi_N1N2CodebookSubsetRestrictiontypeIi)
			return nil
		}
	}
	return nil
}

func (c N1N2CodebookSubsetRestrictiontypeIi) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.N1N2CodebookSubsetRestrictiontypeIi).String(),
		Value: c.N1N2CodebookSubsetRestrictiontypeIi,
	})
}

var _ json.Unmarshaler = (*N1N2CodebookSubsetRestrictiontypeIi)(nil)
