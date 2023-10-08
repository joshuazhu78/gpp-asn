package nrrrcdefinitions

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func Test_NrOfAntennaPortstypeISinglePanel(t *testing.T) {
	c := NewTestNrOfAntennaPortstypeISinglePanel()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled NrOfAntennaPortstypeISinglePanel
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	f, err := ioutil.ReadFile("NrOfAntennaPortstypeISinglePanel.json")
	assert.NoError(t, err)

	d := NrOfAntennaPortstypeISinglePanel{}
	err = json.Unmarshal(f, &d)
	assert.NoError(t, err)
	assert.Equal(t, true, reflect.DeepEqual(*c, d))

}

func NewTestNrOfAntennaPortstypeISinglePanel() *NrOfAntennaPortstypeISinglePanel {
	return &NrOfAntennaPortstypeISinglePanel{
		NrOfAntennaPortstypeISinglePanel: &NrOfAntennaPortstypeISinglePanel_Two{
			Two: &TwonrOfAntennaPorts{
				TwoTxCodebookSubsetRestriction: &asn1.BitString{
					Len:   6,
					Value: []byte{0x3f},
				},
			},
		},
	}
}
