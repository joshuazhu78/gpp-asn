package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s SupportPdschtoHarqackr16OutOfOrderOperationDlr16) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *SupportPdschtoHarqackr16OutOfOrderOperationDlr16) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseSupportPdschtoHarqackr16OutOfOrderOperationDlr16(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseSupportPdschtoHarqackr16OutOfOrderOperationDlr16(s string) (SupportPdschtoHarqackr16OutOfOrderOperationDlr16, error) {
	value, ok := SupportPdschtoHarqackr16OutOfOrderOperationDlr16_value[s]
	if !ok {
		return SupportPdschtoHarqackr16OutOfOrderOperationDlr16(0), fmt.Errorf("%q is not a valid %s", s, "SupportPdschtoHarqackr16OutOfOrderOperationDlr16")
	}
	return SupportPdschtoHarqackr16OutOfOrderOperationDlr16(value), nil
}
