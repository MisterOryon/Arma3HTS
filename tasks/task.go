package tasks

import (
	"errors"
	"fmt"
	"github.com/MisterOryon/Arma3HTS/remote_clients"
	"github.com/MisterOryon/Arma3HTS/utils"
	"github.com/gosuri/uilive"
	"io"
	"os"
)

type TypeTask int

const (
	// UploadFile allow to plan an uploading dir task.
	UploadFile TypeTask = iota

	// DeleteFile allow to plan a deleting dir task.
	DeleteFile

	// CreateDir allow to plan a creation dir task.
	CreateDir
)

// Task allow to plan CreateDir, UploadFile and DeleteFile on remote client.
type Task struct {
	TypeTask   TypeTask
	LocalPath  string
	RemotePath string
	NeedUpload int64
}

// Run allow to apply takes on remote client.
func (t *Task) Run(remoteClient remote_clients.IClient) error {
	handle := func(err error) error {
		return err
	}

	switch {
	case t.TypeTask == CreateDir:
		err := remoteClient.Mkdir(t.RemotePath)
		if err != nil {
			return handle(err)
		}

		fmt.Printf("  - create dir: %s\n", t.RemotePath)

		return nil

	case t.TypeTask == DeleteFile:
		err := remoteClient.Remove(t.RemotePath)
		if err != nil {
			return handle(err)
		}

		fmt.Printf("  - delete file: %s\n", t.RemotePath)

		return nil

	case t.TypeTask == UploadFile:
		source, err := os.Open(t.LocalPath)
		if err != nil {
			return handle(err)
		}
		defer source.Close()

		destination, err := remoteClient.Create(t.RemotePath)
		if err != nil {
			return handle(err)
		}
		defer destination.Close()

		buf := make([]byte, 1048576)

		writer := uilive.New()
		writer.Start()

		for {
			n, err := source.Read(buf)
			if err != nil && err != io.EOF {
				return handle(err)
			}

			if n == 0 {
				break
			}

			if _, err := destination.Write(buf[:n]); err != nil {
				return handle(err)
			}

			t.NeedUpload -= int64(n)
			fmt.Fprintf(writer, "  - upload %s, %s left...\n", t.RemotePath, utils.ConvertSizeToStr(t.NeedUpload))
		}

		writer.Stop()

		return nil

	default:
		return handle(errors.New("unknown task type"))
	}
}
