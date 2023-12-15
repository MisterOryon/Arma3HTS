package remote_clients

import (
	"os"
)

// Provides a level of abstraction to ensure that client files struct have the required functions.
type iFile interface {
	// Write writes len(b) bytes from b to the File. It returns the number of bytes written and an error, if any.
	//Write returns a non-nil error when n != len(b).
	Write(b []byte) (int, error)

	// Close closes the File, rendering it unusable for I/O.
	//On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an ErrClosed error.
	//Close will return an error if it has already been called.
	Close() error
}

// IClient standardized remote clients.
type IClient interface {
	// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.
	Stat(path string) (os.FileInfo, error)

	// ReadDir reads the named directory, returning all its directory entries sorted by filename.
	// If an error occurs reading the directory, ReadDir returns the entries it was able to read before the error, along with the error.
	ReadDir(path string) ([]os.FileInfo, error)

	// Mkdir creates a new directory with the specified name and permission bits (before umask).
	//If there is an error, it will be of type *PathError.
	Mkdir(path string) error

	// Remove removes the named file or directory.
	//If there is an error, it will be of type *PathError.
	Remove(path string) error

	// Create creates or truncates the named file.
	//If the file already exists, it is truncated.
	//If the file does not exist, it is created with mode 0666 (before umask).
	//If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR.
	//If there is an error, it will be of type *PathError.
	Create(path string) (iFile, error)

	// Close allow to disconnect from remote client.
	Close() error
}

type client struct{}
