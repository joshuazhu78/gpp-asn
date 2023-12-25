package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func Test_N1N2CodebookSubsetRestrictionr16TypeIir16(t *testing.T) {
	c := NewTestN1N2CodebookSubsetRestrictionr16TypeIir16()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled N1N2CodebookSubsetRestrictionr16TypeIir16
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestN1N2CodebookSubsetRestrictionr16TypeIir16() *N1N2CodebookSubsetRestrictionr16TypeIir16 {
	return &N1N2CodebookSubsetRestrictionr16TypeIir16{
		N1N2CodebookSubsetRestrictionR16TypeIiR16: &N1N2CodebookSubsetRestrictionr16TypeIir16_FourFour{
			FourFour: &asn1.BitString{
				Len:   139,
				Value: []byte{0x07, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			},
		},
	}
}
