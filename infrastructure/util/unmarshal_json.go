package util

import (
	"encoding/json"
	"io"
)

func ResponseDecoder(r io.Reader) map[string]interface{} {
	var decodedResponse map[string]interface{}
	_ = json.NewDecoder(r).Decode(&decodedResponse)

	return decodedResponse
}
