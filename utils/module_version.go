package utils

import (
	"errors"
	"runtime/debug"
)

// FindMainModuleVersion returns the Arma3HTS version.
func FindMainModuleVersion() (string, error) {
	handle := func(err error) (string, error) {
		return "", err
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return handle(errors.New("failed to read build info"))
	}

	return bi.Main.Version, nil
}
