package remote_clients

import (
	"github.com/pkg/sftp"
	"os"
)

/*
We need adapters in order to map the Create() function of os with the Create() function of the sftp client.
*/

type sftpAdapter struct {
	sc *sftp.Client
}

func (a *sftpAdapter) Stat(path string) (os.FileInfo, error) {
	return a.sc.Stat(path)
}

func (a *sftpAdapter) ReadDir(path string) ([]os.FileInfo, error) {
	return a.sc.ReadDir(path)
}

func (a *sftpAdapter) Mkdir(path string) error {
	return a.sc.Mkdir(path)
}

func (a *sftpAdapter) Remove(path string) error {
	return a.sc.Remove(path)
}

func (a *sftpAdapter) Create(path string) (iFile, error) {
	return a.sc.Create(path)
}
func (a *sftpAdapter) Close() error {
	return a.sc.Close()
}
