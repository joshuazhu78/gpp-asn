package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CsiResourcePeriodicityAndOffset(t *testing.T) {
	c := NewTestCsiResourcePeriodicityAndOffset()
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled CsiResourcePeriodicityAndOffset
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(*c, unmarshaled))

	WriteJsonFile(c, "CsiResourcePeriodicityAndOffset.json")
}

func NewTestCsiResourcePeriodicityAndOffset() *CsiResourcePeriodicityAndOffset {
	return &CsiResourcePeriodicityAndOffset{
		CsiResourcePeriodicityAndOffset: &CsiResourcePeriodicityAndOffset_Slots10{
			Slots10: 0,
		},
	}
}
