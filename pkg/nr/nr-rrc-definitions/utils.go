package nrrrcdefinitions

import (
	"encoding/json"
	"os"
)

func WriteJsonFile(data any, fileName string) {
	file, _ := json.MarshalIndent(data, "", "    ")
	_ = os.WriteFile(fileName, file, 0644)
}

func ReadJsonFile(fileName string, data any) {
	f, _ := os.ReadFile(fileName)
	json.Unmarshal(f, data)
}
