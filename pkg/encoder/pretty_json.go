package encoder

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

// PrettyJSONWithoutIndent formats the given interface and return json string without indentation and tab.
func PrettyJSONWithoutIndent(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)

	_ = encoder.Encode(data)
	return buffer.String()
}

// PrettyJSONWithIndent formats the given interface and return json string with indentation and tab.
func PrettyJSONWithIndent(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	_ = encoder.Encode(data)
	return buffer.String()
}
