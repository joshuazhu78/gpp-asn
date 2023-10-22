package nrrrcdefinitions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReportIntervalContainer struct {
	ReportInterval ReportInterval
}

func Test_ReportInterval(t *testing.T) {
	c := ReportIntervalContainer{
		ReportInterval: ReportInterval(0),
	}
	data, err := json.Marshal(c)
	assert.NoError(t, err)

	var unmarshaled ReportIntervalContainer
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, true, reflect.DeepEqual(c, unmarshaled))

	WriteJsonFile(c, "ReportInterval.json")
	d := ReportIntervalContainer{}
	ReadJsonFile("ReportInterval.json", &d)
	assert.NoError(t, err)
	assert.Equal(t, true, reflect.DeepEqual(c, d))

}
