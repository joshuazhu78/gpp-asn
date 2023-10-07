package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CodebookTypeCodebookConfig(t *testing.T) {
	c := NewTestCodebookTypeCodebookConfig()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CodebookTypeCodebookConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestCodebookTypeCodebookConfig() *CodebookTypeCodebookConfig {
	return &CodebookTypeCodebookConfig{
		CodebookTypeCodebookConfig: &CodebookTypeCodebookConfig_Type2{
			Type2: &Type2CodebookType{
				SubType:           NewTestSubTypetype2(),
				PhaseAlphabetSize: PhaseAlphabetSizetype2_PHASE_ALPHABET_SIZETYPE2_N4,
				SubbandAmplitude:  false,
			},
		},
	}
}
