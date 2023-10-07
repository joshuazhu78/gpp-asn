package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CodebookTyper16CodebookConfigr16(t *testing.T) {
	c := NewTestCodebookTyper16CodebookConfigr16()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CodebookTyper16CodebookConfigr16
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestCodebookTyper16CodebookConfigr16() *CodebookTyper16CodebookConfigr16 {
	return &CodebookTyper16CodebookConfigr16{
		CodebookTypeR16CodebookConfigR16: &CodebookTyper16CodebookConfigr16_Type2R16{
			Type2R16: &Type2R16CodebookTyper16{
				SubType:                             NewTestSubTypetype2R16(),
				NumberOfPmiSubbandsPerCqiSubbandR16: 2,
				ParamCombinationR16:                 1,
			},
		},
	}
}
