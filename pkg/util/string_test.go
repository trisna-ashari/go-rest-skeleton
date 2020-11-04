package util_test

import (
	"go-rest-skeleton/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	sliceStr := []string{"en", "id"}
	assert.EqualValues(t, util.SliceContains(sliceStr, "en"), true)
	assert.EqualValues(t, util.SliceContains(sliceStr, "es"), false)
}

func TestSentenceCase(t *testing.T) {
	message := "Not Found"
	assert.Equal(t, "Not found", util.SentenceCase(message))
}
