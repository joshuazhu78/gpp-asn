package nrrrcdefinitions

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON must be a *value receiver* to ensure that a Suit on a parent object
// does not have to be a pointer in order to have it correctly marshaled.
func (s ReportSlotConfigv1530SemiPersistentOnPuschv1530) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *ReportSlotConfigv1530SemiPersistentOnPuschv1530) UnmarshalJSON(data []byte) (err error) {
	var reportInterval string
	if err := json.Unmarshal(data, &reportInterval); err != nil {
		return err
	}
	if *s, err = ParseReportSlotConfigv1530SemiPersistentOnPuschv1530(reportInterval); err != nil {
		return err
	}
	return nil
}

func ParseReportSlotConfigv1530SemiPersistentOnPuschv1530(s string) (ReportSlotConfigv1530SemiPersistentOnPuschv1530, error) {
	value, ok := ReportSlotConfigv1530SemiPersistentOnPuschv1530_value[s]
	if !ok {
		return ReportSlotConfigv1530SemiPersistentOnPuschv1530(0), fmt.Errorf("%q is not a valid %s", s, "ReportSlotConfigv1530SemiPersistentOnPuschv1530")
	}
	return ReportSlotConfigv1530SemiPersistentOnPuschv1530(value), nil
}
