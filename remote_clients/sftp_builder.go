package remote_clients

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sftpBuilder struct {
	host     string
	port     int
	user     string
	password string
}

func (b *sftpBuilder) SetHost(host string) {
	b.host = host
}

func (b *sftpBuilder) SetPort(port int) {
	b.port = port
}

func (b *sftpBuilder) SetUser(user string) {
	b.user = user
}

func (b *sftpBuilder) SetPassword(password string) {
	b.password = password
}

func (b *sftpBuilder) GetClient() (IClient, error) {
	handle := func(err error) (IClient, error) {
		return nil, err
	}

	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(b.password))

	config := ssh.ClientConfig{
		User:            b.user,
		Auth:            auths,
		HostKeyCallback: hostKeyCallback,
	}

	addr := fmt.Sprintf("%s:%d", b.host, b.port)

	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return handle(err)
	}

	sc, err := sftp.NewClient(conn, sftp.UseConcurrentWrites(true))
	if err != nil {
		return handle(err)
	}

	return &sftpAdapter{sc: sc}, nil
}
