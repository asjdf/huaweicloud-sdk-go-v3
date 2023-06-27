package utils

import (
	"encoding/json"
)

func Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func Unmarshal(data []byte, i interface{}) error {
	return json.Unmarshal(data, i)
}
