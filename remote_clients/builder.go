package remote_clients

import "errors"

// IBuilder standardized builder for remote clients.
type IBuilder interface {
	// SetHost allow set the host of remote client.
	SetHost(host string)

	// SetPort allow set the port of remote client.
	SetPort(port int)

	// SetUser allow set the user for connecting to remote client.
	SetUser(user string)

	// SetPassword allow set the password for connecting to remote client.
	SetPassword(password string)

	// GetClient returns standardized sftp client.
	GetClient() (IClient, error)
}

// NewBuilder returns a standardized builder for remote clients.
func NewBuilder(typeBuilder string) (IBuilder, error) {
	handle := func(err error) (IBuilder, error) {
		return nil, err
	}

	switch {
	case typeBuilder == "sftp":
		return &sftpBuilder{}, nil
	default:
		return handle(errors.New("unknown type builder"))
	}
}
