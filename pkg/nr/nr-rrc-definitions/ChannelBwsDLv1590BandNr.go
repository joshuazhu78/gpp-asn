package nrrrcdefinitions

import (
	"encoding/json"
	reflect "reflect"
)

// We need to register all known message types here to be able to unmarshal them to the correct interface type.
var ChannelBwsDlV1590BandNrknownImplementations = []isChannelBwsDLv1590BandNr_ChannelBwsDlV1590BandNr{
	&ChannelBwsDLv1590BandNr_Fr1{},
	&ChannelBwsDLv1590BandNr_Fr2{},
}

func (c *ChannelBwsDLv1590BandNr) UnmarshalJSON(bytes []byte) error {
	var data struct {
		Type  string
		Value json.RawMessage
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	for _, knownImplementation := range ChannelBwsDlV1590BandNrknownImplementations {
		knownType := reflect.TypeOf(knownImplementation)
		if knownType.String() == data.Type {
			// Create a new pointer to a value of the concrete message type
			target := reflect.New(knownType)

			// Unmarshal the data to an interface to the concrete value (which will act as a pointer, don't ask why)
			if err := json.Unmarshal(data.Value, target.Interface()); err != nil {
				return err
			}
			// Now we get the element value of the target and convert it to the interface type (this is to get rid of a pointer type instead of a plain struct value)
			c.ChannelBwsDlV1590BandNr = target.Elem().Interface().(isChannelBwsDLv1590BandNr_ChannelBwsDlV1590BandNr)
			return nil
		}
	}
	return nil
}

func (c ChannelBwsDLv1590BandNr) MarshalJSON() ([]byte, error) {
	// Marshal to type and actual data to handle unmarshaling to specific interface type
	return json.Marshal(struct {
		Type  string
		Value any
	}{
		Type:  reflect.TypeOf(c.ChannelBwsDlV1590BandNr).String(),
		Value: c.ChannelBwsDlV1590BandNr,
	})
}

var _ json.Unmarshaler = (*ChannelBwsDLv1590BandNr)(nil)
