package util

import (
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Join(filepath.Dir(b), "../..")
	}
	return empty
}
