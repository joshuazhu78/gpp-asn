package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func Test_SubTypetype2(t *testing.T) {
	c := NewTestSubTypetype2()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled SubTypetype2
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestSubTypetype2() *SubTypetype2 {
	return &SubTypetype2{
		SubTypetype2: &SubTypetype2_TypeIi{
			TypeIi: &TypeIisubType{
				N1N2CodebookSubsetRestriction: NewTestN1N2CodebookSubsetRestrictiontypeIi(),
				TypeIiRiRestriction: &asn1.BitString{
					Len:   2,
					Value: []byte{0x03},
				},
			},
		},
	}
}
