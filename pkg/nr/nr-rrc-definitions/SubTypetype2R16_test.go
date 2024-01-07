package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func Test_SubTypetype2R16(t *testing.T) {
	c := NewTestSubTypetype2R16()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled SubTypetype2R16
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestSubTypetype2R16() *SubTypetype2R16 {
	return &SubTypetype2R16{
		SubTypetype2R16: &SubTypetype2R16_TypeIiR16{
			TypeIiR16: &TypeIir16SubType{
				N1N2CodebookSubsetRestrictionR16: NewTestN1N2CodebookSubsetRestrictionr16TypeIir16(),
				TypeIiRiRestrictionR16: &asn1.BitString{
					Len:   4,
					Value: []byte{0x02},
				},
			},
		},
	}
}
