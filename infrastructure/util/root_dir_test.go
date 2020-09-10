package util_test

import (
	"go-rest-skeleton/infrastructure/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootDir(t *testing.T) {
	rootDir := util.RootDir()
	assert.NotEmpty(t, rootDir)
}
