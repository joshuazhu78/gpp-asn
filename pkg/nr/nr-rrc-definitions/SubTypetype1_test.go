package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func Test_SubTypetype1(t *testing.T) {
	c := NewTestSubTypetype1(true)
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled SubTypetype1
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))
}

func NewTestSubTypetype1(twoPort bool) *SubTypetype1 {
	if twoPort {
		return &SubTypetype1{
			SubTypetype1: &SubTypetype1_TypeISinglePanel{
				TypeISinglePanel: &TypeISinglePanelsubType{
					NrOfAntennaPorts: &NrOfAntennaPortstypeISinglePanel{
						NrOfAntennaPortstypeISinglePanel: &NrOfAntennaPortstypeISinglePanel_Two{
							Two: &TwonrOfAntennaPorts{
								TwoTxCodebookSubsetRestriction: &asn1.BitString{
									Len:   6,
									Value: []byte{0x3f},
								},
							},
						},
					},
					TypeISinglePanelRiRestriction: &asn1.BitString{
						Len:   2,
						Value: []byte{0x03},
					},
				},
			},
		}
	}
	return &SubTypetype1{
		SubTypetype1: &SubTypetype1_TypeISinglePanel{
			TypeISinglePanel: &TypeISinglePanelsubType{
				NrOfAntennaPorts: &NrOfAntennaPortstypeISinglePanel{
					NrOfAntennaPortstypeISinglePanel: &NrOfAntennaPortstypeISinglePanel_MoreThanTwo{
						MoreThanTwo: &MoreThanTwonrOfAntennaPorts{
							N1N2: &N1N2MoreThanTwo{
								N1N2MoreThanTwo: &N1N2MoreThanTwo_FourFourTypeISinglePanelRestriction{
									FourFourTypeISinglePanelRestriction: &asn1.BitString{
										Len: 256,
										Value: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
											0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
											0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
											0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
									},
								},
							},
							TypeISinglePanelCodebookSubsetRestrictionI2: &asn1.BitString{
								Len:   16,
								Value: []byte{0xff, 0xff},
							},
						},
					},
				},
				TypeISinglePanelRiRestriction: &asn1.BitString{
					Len:   2,
					Value: []byte{0x03},
				},
			},
		},
	}
}
