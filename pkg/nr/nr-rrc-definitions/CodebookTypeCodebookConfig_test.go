package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CodebookTypeCodebookConfigType2(t *testing.T) {
	c := NewTestCodebookTypeCodebookConfigType2()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CodebookTypeCodebookConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestCodebookTypeCodebookConfigType1(twoPort bool) *CodebookTypeCodebookConfig {
	return &CodebookTypeCodebookConfig{
		CodebookTypeCodebookConfig: &CodebookTypeCodebookConfig_Type1{
			Type1: &Type1CodebookType{
				SubType:      NewTestSubTypetype1(twoPort),
				CodebookMode: 1,
			},
		},
	}
}

func NewTestCodebookTypeCodebookConfigType2() *CodebookTypeCodebookConfig {
	return &CodebookTypeCodebookConfig{
		CodebookTypeCodebookConfig: &CodebookTypeCodebookConfig_Type2{
			Type2: &Type2CodebookType{
				SubType:           NewTestSubTypetype2(),
				PhaseAlphabetSize: PhaseAlphabetSizetype2_PHASE_ALPHABET_SIZETYPE2_N8,
				SubbandAmplitude:  false,
				NumberOfBeams:     2,
			},
		},
	}
}
