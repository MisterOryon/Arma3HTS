package mods

import "fmt"

type CheckModErrorType int

const (
	// DirNotFound is chosen when a mod folder has not present on the computer/server .
	DirNotFound CheckModErrorType = iota

	// IsNotDir is chosen when a file has the name a mod folder should have.
	IsNotDir

	// DirUseless is chosen when a mod folder is present on the computer/server but no in th preset.
	DirUseless

	// Unknown is chosen when the error is not DirNotFound, IsNotDir or DirUseless.
	Unknown
)

// String convert CheckModErrorType to string.
func (t CheckModErrorType) String() string {
	errorTypes := []string{"folder not found", "is not a folder", "folder is useless"}

	if t >= DirNotFound || t < Unknown {
		return errorTypes[t]
	}

	return "unknown error"
}

// CheckModError custom error redound by check mods proses.
type CheckModError struct {
	_type CheckModErrorType
	path  string
	err   error
}

// Error return a custom error.
func (m *CheckModError) Error() string {
	switch {
	case m._type == DirNotFound:
		return fmt.Sprintf("Unable to find folder '%s'", m.path)

	case m._type == IsNotDir:
		return fmt.Sprintf("'%s' is not a folder", m.path)

	case m._type == DirUseless:
		return fmt.Sprintf("'%s' is useless", m.path)

	default:
		return fmt.Sprintf("An error has occurred!")
	}
}

// GetErr returns a debug error.
func (m *CheckModError) GetErr() error {
	return m.err
}
