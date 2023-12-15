package remote_clients

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"net"
	"os"
	"path"
	"runtime"
)

// Returns the known_hosts path based on the OS the cli is running on.
func knownHostsPath() string {
	if runtime.GOOS == "windows" {
		return path.Join(os.Getenv("USERPROFILE"), ".ssh/known_hosts")
	} else {
		return path.Join(os.Getenv("HOME"), ".ssh/known_hosts")
	}
}

// check the known_hosts to validate the connection.
var hostKeyCallback = ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error {
	handle := func(err error) error {
		return err
	}

	kh, err := knownhosts.New(knownHostsPath())
	if err != nil {
		return handle(errors.New("unable to read known_hosts file"))
	}

	// Check remote client credentials.
	err = kh(hostname, remote, key)
	var keyErr *knownhosts.KeyError

	switch {
	case errors.As(err, &keyErr) && len(keyErr.Want) > 0:
		// Call when the credentials of the known_hosts file do not match with the remote client,
		// which may mean a man-in-the-middle attack.

		return handle(errors.New("remote host has changed, which may mean a man-in-the-middle attack"))

	case errors.As(err, &keyErr) && len(keyErr.Want) == 0:
		// Call on first connection to the remote client to add these credentials to known_hosts file.

		known_hosts, err := os.OpenFile(knownHostsPath(), os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return handle(err)
		}
		defer known_hosts.Close()

		addresses := []string{knownhosts.Normalize(hostname), knownhosts.Normalize(remote.String())}

		_, err = known_hosts.Write([]byte(knownhosts.Line(addresses, key) + "\n"))
		return handle(err)
	default:
		return handle(err)
	}
})
