package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CsiMeasConfigServingCellConfig(t *testing.T) {
	c := NewTestCsiMeasConfigServingCellConfig()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiMeasConfigServingCellConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestCsiMeasConfigServingCellConfig() *CsiMeasConfigServingCellConfig {
	return &CsiMeasConfigServingCellConfig{
		CsiMeasConfigServingCellConfig: &CsiMeasConfigServingCellConfig_Setup{
			Setup: &CsiMeasConfig{
				CsiReportConfigToAddModList: []*CsiReportConfig{
					{
						CodebookConfig: &CodebookConfig{
							CodebookType: NewTestCodebookTypeCodebookConfig(),
						},
						CodebookConfigR16: &CodebookConfigr16{
							CodebookTypeR16: NewTestCodebookTyper16CodebookConfigr16(),
						},
					},
				},
			},
		},
	}
}
