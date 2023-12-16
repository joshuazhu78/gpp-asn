package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CellGroupConfig20GHz(t *testing.T) {
	c := NewTestCellGroupConfig(1, 0.6, DlUltransmissionPeriodicityTdduldlpattern_DL_ULTRANSMISSION_PERIODICITY_TDDULDLPATTERN_MS2P5, 2112499, SubcarrierSpacing_SUBCARRIER_SPACING_K_HZ120, 17875, false)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CellGroupConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CellGroupConfig20GHz.json")
}

func Test_CellGroupConfig6GHzType2(t *testing.T) {
	c := NewTestCellGroupConfig(1, 0.6, DlUltransmissionPeriodicityTdduldlpattern_DL_ULTRANSMISSION_PERIODICITY_TDDULDLPATTERN_MS2P5, 800000, SubcarrierSpacing_SUBCARRIER_SPACING_K_HZ30, 13750, false)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CellGroupConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CellGroupConfig6GHzType2.json")
}

func Test_CellGroupConfig6GHzR16Type2(t *testing.T) {
	c := NewTestCellGroupConfig(1, 0.6, DlUltransmissionPeriodicityTdduldlpattern_DL_ULTRANSMISSION_PERIODICITY_TDDULDLPATTERN_MS2P5, 800000, SubcarrierSpacing_SUBCARRIER_SPACING_K_HZ30, 13750, true)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CellGroupConfig
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CellGroupConfig6GHzR16Type2.json")
}

func NewTestCellGroupConfig(physCellId int, downlinkRatio float64, dlUltransmissionPeriodicityTdduldlpattern DlUltransmissionPeriodicityTdduldlpattern, arfcnValueNr int, subcarrierSpacing SubcarrierSpacing, locationAndBandwidth int, r16 bool) *CellGroupConfig {
	slotsPerFrame := 10 * (1 << subcarrierSpacing)
	slotsPerPeriod := int(float64(slotsPerFrame) * (DlUltransmissionPeriodicityTdduldlpatternMap[dlUltransmissionPeriodicityTdduldlpattern] / 10.0))
	nrofDownlinkSlots := int(float64(slotsPerPeriod) * downlinkRatio)
	return &CellGroupConfig{
		CellGroupId: &CellGroupId{
			Value: 1,
		},
		SpCellConfig: &SpCellConfig{
			ReconfigurationWithSync: &ReconfigurationWithSync{
				SpCellConfigCommon: &ServingCellConfigCommon{
					PhysCellId: &PhysCellId{
						Value: int32(physCellId),
					},
					DownlinkConfigCommon: &DownlinkConfigCommon{
						FrequencyInfoDl: &FrequencyInfoDl{
							AbsoluteFrequencyPointA: &ArfcnValueNr{
								Value: int32(arfcnValueNr),
							},
							ScsSpecificCarrierList: []*ScsSpecificCarrier{
								{
									SubcarrierSpacing: subcarrierSpacing,
								},
							},
						},
						InitialDownlinkBwp: &BwpDownlinkCommon{
							GenericParameters: &Bwp{
								SubcarrierSpacing:    subcarrierSpacing,
								LocationAndBandwidth: int32(locationAndBandwidth),
							},
						},
					},
					TddUlDlConfigurationCommon: &TddULDLConfigCommon{
						ReferenceSubcarrierSpacing: subcarrierSpacing,
						Pattern1: &TddULDLPattern{
							DlUlTransmissionPeriodicity: DlUltransmissionPeriodicityTdduldlpattern_DL_ULTRANSMISSION_PERIODICITY_TDDULDLPATTERN_MS2P5,
							NrofDownlinkSlots:           int32(nrofDownlinkSlots),
							NrofUplinkSlots:             int32(slotsPerPeriod - nrofDownlinkSlots - 1),
							NrofDownlinkSymbols:         4,
							NrofUplinkSymbols:           10,
						},
					},
				},
			},
			SpCellConfigDedicated: NewTestServingCellConfig(subcarrierSpacing, locationAndBandwidth, r16),
		},
	}
}

func NewTestServingCellConfig(subcarrierSpacing SubcarrierSpacing, locationAndBandwidth int, r16 bool) *ServingCellConfig {
	//slotsPerFrame := 20 * (1 << subcarrierSpacing)
	//downlinkRatio := 0.8
	//downlinkSlots := int(downlinkRatio * float64(slotsPerFrame))
	servingCellConfig := &ServingCellConfig{
		DownlinkBwpToAddModList: []*BwpDownlink{
			{
				BwpId: &BwpId{
					Value: 0,
				},
				BwpCommon: &BwpDownlinkCommon{
					GenericParameters: &Bwp{
						SubcarrierSpacing:    subcarrierSpacing,
						LocationAndBandwidth: int32(locationAndBandwidth),
					},
				},
			},
		},
		UplinkConfig: &UplinkConfig{
			UplinkBwpToAddModList: []*BwpUplink{
				{
					BwpId: &BwpId{
						Value: 0,
					},
					BwpCommon: &BwpUplinkCommon{
						GenericParameters: &Bwp{
							SubcarrierSpacing:    subcarrierSpacing,
							LocationAndBandwidth: int32(locationAndBandwidth),
						},
					},
				},
			},
		},
	}
	if subcarrierSpacing == SubcarrierSpacing_SUBCARRIER_SPACING_K_HZ120 {
		servingCellConfig.CsiMeasConfig = NewTestCsiMeasConfigServingCellConfigCodebookConfigType1(true)
	} else {
		if r16 {
			servingCellConfig.CsiMeasConfig = NewTestCsiMeasConfigServingCellConfigCodebookConfigr16()
		} else {
			servingCellConfig.CsiMeasConfig = NewTestCsiMeasConfigServingCellConfigCodebookConfigType2()
		}
	}

	/*servingCellConfig.TddUlDlConfigurationDedicated = &TddULDLConfigDedicated{
		SlotSpecificConfigurationsToAddModList: make([]*TddULDLSlotConfig, slotsPerFrame),
	}
	slotIdx := 0
	for ; slotIdx < downlinkSlots; slotIdx++ {
		servingCellConfig.TddUlDlConfigurationDedicated.SlotSpecificConfigurationsToAddModList[slotIdx] = &TddULDLSlotConfig{
			SlotIndex: &TddULDLSlotIndex{
				Value: int32(slotIdx),
			},
			Symbols: &SymbolsTddULDLSlotConfig{
				SymbolsTddUlDlSlotConfig: &SymbolsTddULDLSlotConfig_AllDownlink{},
			},
		}
	}
	var nrofDownlinkSymbols int32 = 3
	var nrofUplinkSymbols int32 = 14 - nrofDownlinkSymbols
	for ; slotIdx < slotsPerFrame; slotIdx++ {
		servingCellConfig.TddUlDlConfigurationDedicated.SlotSpecificConfigurationsToAddModList[slotIdx] = &TddULDLSlotConfig{
			SlotIndex: &TddULDLSlotIndex{
				Value: int32(slotIdx),
			},
			Symbols: &SymbolsTddULDLSlotConfig{
				SymbolsTddUlDlSlotConfig: &SymbolsTddULDLSlotConfig_Explicit{
					Explicit: &Explicitsymbols{
						NrofDownlinkSymbols: &nrofDownlinkSymbols,
						NrofUplinkSymbols:   &nrofUplinkSymbols,
					},
				},
			},
		}
	}*/

	return servingCellConfig
}
