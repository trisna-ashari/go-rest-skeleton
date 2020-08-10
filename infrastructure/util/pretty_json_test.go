package util_test

import (
	"go-rest-skeleton/infrastructure/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JSONString struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func TestPrettyJSON(t *testing.T) {
	data := JSONString{
		Code:    200,
		Data:    nil,
		Message: "OK",
	}
	json := util.PrettyJSON(data)
	expectedJSON := "{\n\t\"code\": 200,\n\t\"data\": null,\n\t\"message\": \"OK\"\n}\n"
	assert.Equal(t, expectedJSON, json)
}
