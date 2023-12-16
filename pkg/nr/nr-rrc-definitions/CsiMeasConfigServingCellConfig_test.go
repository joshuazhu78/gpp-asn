package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CsiMeasConfigServingCellConfigConfigCodebookConfigType1Two(t *testing.T) {
	c := NewTestCsiMeasConfigServingCellConfigCodebookConfigType1(true)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiMeasConfigServingCellConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CsiMeasConfigServingCellConfigCodebookConfigType1Two.json")
	var d CsiMeasConfigServingCellConfig
	ReadJsonFile("CsiMeasConfigServingCellConfigCodebookConfigType1Two.json", &d)
	assert.Equal(t, true, reflect.DeepEqual(*c, d))
}

func Test_CsiMeasConfigServingCellConfigConfigCodebookConfigType1FourFour(t *testing.T) {
	c := NewTestCsiMeasConfigServingCellConfigCodebookConfigType1(false)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiMeasConfigServingCellConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CsiMeasConfigServingCellConfigCodebookConfigType1FourFour.json")
	var d CsiMeasConfigServingCellConfig
	ReadJsonFile("CsiMeasConfigServingCellConfigCodebookConfigType1FourFour.json", &d)
	assert.Equal(t, true, reflect.DeepEqual(*c, d))
}

func Test_CsiMeasConfigServingCellConfigConfigCodebookConfigType2(t *testing.T) {
	c := NewTestCsiMeasConfigServingCellConfigCodebookConfigType2()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiMeasConfigServingCellConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CsiMeasConfigServingCellConfigCodebookConfigType2.json")
	var d CsiMeasConfigServingCellConfig
	ReadJsonFile("CsiMeasConfigServingCellConfigCodebookConfigType2.json", &d)
	assert.Equal(t, true, reflect.DeepEqual(*c, d))
}

func Test_CsiMeasConfigServingCellConfigConfigCodebookConfigr16(t *testing.T) {
	c := NewTestCsiMeasConfigServingCellConfigCodebookConfigr16()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiMeasConfigServingCellConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CsiMeasConfigServingCellConfigCodebookConfigr16.json")
	var d CsiMeasConfigServingCellConfig
	ReadJsonFile("CsiMeasConfigServingCellConfigCodebookConfigr16.json", &d)
	assert.Equal(t, true, reflect.DeepEqual(*c, d))
}

func NewTestCsiMeasConfigServingCellConfigCodebookConfigType1(twoPort bool) *CsiMeasConfigServingCellConfig {
	var PmiFormatIndicator PmiFormatIndicatorreportFreqConfiguration = PmiFormatIndicatorreportFreqConfiguration_PMI_FORMAT_INDICATORREPORT_FREQ_CONFIGURATION_WIDEBAND_PMI
	var NrofReportedRsdisabled = NrofReportedRsdisabled_NROF_REPORTED_RSDISABLED_N1
	return &CsiMeasConfigServingCellConfig{
		CsiMeasConfigServingCellConfig: &CsiMeasConfigServingCellConfig_Setup{
			Setup: &CsiMeasConfig{
				CsiReportConfigToAddModList: []*CsiReportConfig{
					{
						SubbandSize: SubbandSizeCsireportConfig_SUBBAND_SIZE_CSIREPORT_CONFIG_VALUE2,
						ReportFreqConfiguration: &ReportFreqConfigurationCsiReportConfig{
							PmiFormatIndicator: &PmiFormatIndicator,
						},
						CodebookConfig: &CodebookConfig{
							CodebookType: NewTestCodebookTypeCodebookConfigType1(twoPort),
						},
						ReportQuantity: &ReportQuantityCsiReportConfig{
							ReportQuantityCsiReportConfig: &ReportQuantityCsiReportConfig_CriRiPmiCqi{},
						},
						GroupBasedBeamReporting: &GroupBasedBeamReportingCsiReportConfig{
							GroupBasedBeamReportingCsiReportConfig: &GroupBasedBeamReportingCsiReportConfig_Disabled{
								Disabled: &DisabledgroupBasedBeamReporting{
									NrofReportedRs: &NrofReportedRsdisabled,
								},
							},
						},
					},
				},
			},
		},
	}
}

func NewTestCsiMeasConfigServingCellConfigCodebookConfigType2() *CsiMeasConfigServingCellConfig {
	var PmiFormatIndicator PmiFormatIndicatorreportFreqConfiguration = PmiFormatIndicatorreportFreqConfiguration_PMI_FORMAT_INDICATORREPORT_FREQ_CONFIGURATION_WIDEBAND_PMI
	return &CsiMeasConfigServingCellConfig{
		CsiMeasConfigServingCellConfig: &CsiMeasConfigServingCellConfig_Setup{
			Setup: &CsiMeasConfig{
				CsiReportConfigToAddModList: []*CsiReportConfig{
					{
						SubbandSize: SubbandSizeCsireportConfig_SUBBAND_SIZE_CSIREPORT_CONFIG_VALUE2,
						ReportFreqConfiguration: &ReportFreqConfigurationCsiReportConfig{
							PmiFormatIndicator: &PmiFormatIndicator,
						},
						CodebookConfig: &CodebookConfig{
							CodebookType: NewTestCodebookTypeCodebookConfigType2(),
						},
					},
				},
			},
		},
	}
}

func NewTestCsiMeasConfigServingCellConfigCodebookConfigr16() *CsiMeasConfigServingCellConfig {
	var PmiFormatIndicator PmiFormatIndicatorreportFreqConfiguration = PmiFormatIndicatorreportFreqConfiguration_PMI_FORMAT_INDICATORREPORT_FREQ_CONFIGURATION_WIDEBAND_PMI
	return &CsiMeasConfigServingCellConfig{
		CsiMeasConfigServingCellConfig: &CsiMeasConfigServingCellConfig_Setup{
			Setup: &CsiMeasConfig{
				CsiReportConfigToAddModList: []*CsiReportConfig{
					{
						SubbandSize: SubbandSizeCsireportConfig_SUBBAND_SIZE_CSIREPORT_CONFIG_VALUE2,
						ReportFreqConfiguration: &ReportFreqConfigurationCsiReportConfig{
							PmiFormatIndicator: &PmiFormatIndicator,
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
