package util

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJSON(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	_ = encoder.Encode(data)
	return buffer.String()
}
