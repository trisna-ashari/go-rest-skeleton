package json_formatter_test

import (
	"go-rest-skeleton/pkg/json_formatter"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JSONString struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func TestPrettyJSONWithoutIndent(t *testing.T) {
	data := JSONString{
		Code:    200,
		Data:    nil,
		Message: "OK",
	}
	json := json_formatter.PrettyJSONWithoutIndent(data)
	expectedJSON := "{\"code\":200,\"data\":null,\"message\":\"OK\"}\n"
	assert.Equal(t, expectedJSON, json)
}

func TestPrettyJSONWithIndent(t *testing.T) {
	data := JSONString{
		Code:    200,
		Data:    nil,
		Message: "OK",
	}
	json := json_formatter.PrettyJSONWithIndent(data)
	expectedJSON := "{\n\t\"code\": 200,\n\t\"data\": null,\n\t\"message\": \"OK\"\n}\n"
	assert.Equal(t, expectedJSON, json)
}
