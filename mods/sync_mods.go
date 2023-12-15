package mods

import (
	"fmt"
	"github.com/MisterOryon/Arma3HTS/remote_clients"
	"github.com/MisterOryon/Arma3HTS/tasks"
	"github.com/MisterOryon/Arma3HTS/utils"
	"github.com/gosuri/uilive"
	"os"
	"path"
	"strconv"
	"strings"
)

var taskQueue *tasks.TaskQueue

// SyncMods allows you to synchronize the mods of a local machine to a remote machine using an Arma 3 preset.
// Once the analysis is completed, this function returns a waiting queue of tasks to be carried out.
func SyncMods(workshopPath string, remotePath string, listMods *[]mod, client remote_clients.IClient) (*tasks.TaskQueue, error) {
	// Clear last queue.
	taskQueue = &tasks.TaskQueue{}

	handle := func(err error) (*tasks.TaskQueue, error) {
		return nil, err
	}

	remoteDir, err := client.ReadDir(remotePath)
	if err != nil {
		return handle(err)
	}

	writer := uilive.New()
	writer.Start()

	// Scans remote client folders to find mod folders.
	for i, mod := range remoteDir {
		// Don't look folders which are not mod folders.
		if !strings.HasPrefix(mod.Name(), "@") {
			continue
		}

		fmt.Fprintf(writer, "Scanning folders... (%d/%d)\n", i, len(remoteDir))

		modPath := path.Join(remotePath, mod.Name())

		// Delete the @ at the beginning of the mod name.
		modName, err := strconv.Atoi(mod.Name()[1:])
		if err != nil {
			return handle(err)
		}

		index := IndexOfModsArray(listMods, modName)

		switch {
		case index == -1:
			// If the folder was not found in the list of mods it is because it is an additional folder,
			// which must be deleted because it is no longer used by the preset.

			err := needDeleteDir(modPath, client)
			if err != nil {
				return handle(err)
			}

		default:
			// Start scanning files in mod folder.

			localModPath := path.Join(workshopPath, "@"+(*listMods)[index].GetName())
			remoteModPath := path.Join(remotePath, mod.Name())

			err := syncFolder(localModPath, remoteModPath, client)
			if err != nil {
				return handle(err)
			}

			*listMods = append((*listMods)[:index], (*listMods)[index+1:]...)
		}
	}

	// After analyzing of folder, only missed folders remain in listMods since they are not present on the remote client,
	// it must be necessary to uploading from the server.
	for _, modsMiss := range *listMods {
		localModPath := path.Join(workshopPath, "@"+modsMiss.GetName())
		remoteModPath := path.Join(remotePath, "@"+strconv.Itoa(modsMiss.GetId()))

		err := syncFolder(localModPath, remoteModPath, client)
		if err != nil {
			return handle(err)
		}
	}

	fmt.Fprintln(writer, "Scanning mod folders... (done)")
	writer.Stop()

	return taskQueue, nil
}

// Compares the contents of a folder between the local machine and the remote machine.
func syncFolder(localDirPath string, remoteDirPath string, client remote_clients.IClient) error {
	handle := func(err error) error {
		return err
	}

	localFiles, err := extractFileInfoOfDir(localDirPath)
	if err != nil {
		return handle(err)
	}

	// Pre-test that the folder exists using the State function,
	// because the ReadDir function does not return os.IsNotExist errors in case of a non-existent folder.
	_, err = client.Stat(remoteDirPath)
	switch {
	case err != nil && os.IsNotExist(err):
		return needUploadDir(localDirPath, remoteDirPath)

	case err != nil:
		return handle(err)
	}

	remoteFiles, err := client.ReadDir(remoteDirPath)
	if err != nil {
		return handle(err)
	}

	for _, localFile := range localFiles {
		index := IndexOfFileInfoArray(&remoteFiles, localFile.Name())

		switch {
		case index == -1 && localFile.IsDir():
			// Call when the folder is not present on the remote server.

			newLocalDirPath := path.Join(localDirPath, localFile.Name())
			newRemoteDirPath := path.Join(remoteDirPath, utils.NormalizedFileName(localFile.Name()))

			err := needUploadDir(newLocalDirPath, newRemoteDirPath)
			if err != nil {
				return handle(err)
			}

		case index == -1 && !localFile.IsDir():
			// Call when the file is not present on the remote server.

			newLocalFilePath := path.Join(localDirPath, localFile.Name())
			newRemoteFilePath := path.Join(remoteDirPath, utils.NormalizedFileName(localFile.Name()))

			taskQueue.AddTask(&tasks.Task{
				TypeTask:   tasks.UploadFile,
				LocalPath:  newLocalFilePath,
				RemotePath: newRemoteFilePath,
				NeedUpload: localFile.Size(),
			})

		case index != -1 && localFile.IsDir():
			// Call to continue scanning files in the new-found folder.

			newLocalDirPath := path.Join(localDirPath, localFile.Name())
			newRemoteDirPath := path.Join(remoteDirPath, remoteFiles[index].Name())

			err := syncFolder(newLocalDirPath, newRemoteDirPath, client)
			if err != nil {
				return handle(err)
			}

			// Remove the current file from the list.
			remoteFiles = append(remoteFiles[:index], remoteFiles[index+1:]...)

		case index != -1 && !localFile.IsDir():
			// Call to check that a file is up-to-date.

			// TODO fix this
			if localFile.Size() != remoteFiles[index].Size() {
				newLocalFilePath := path.Join(localDirPath, localFile.Name())
				newRemoteFilePath := path.Join(remoteDirPath, remoteFiles[index].Name())

				taskQueue.AddTask(&tasks.Task{
					TypeTask:   tasks.UploadFile,
					LocalPath:  newLocalFilePath,
					RemotePath: newRemoteFilePath,
					NeedUpload: localFile.Size(),
				})
			}

			// Remove the current file from the list.
			remoteFiles = append(remoteFiles[:index], remoteFiles[index+1:]...)
		}
	}

	// After analyzing of folder, only unnecessary files remain in remoteFiles since they are no longer present on the host,
	// must be deleted from the server.
	for _, uselessRemoteFile := range remoteFiles {
		newRemotePath := path.Join(remoteDirPath, uselessRemoteFile.Name())

		if uselessRemoteFile.IsDir() {
			err := needDeleteDir(newRemotePath, client)
			if err != nil {
				return handle(err)
			}
		} else {
			taskQueue.AddTask(&tasks.Task{
				TypeTask:   tasks.DeleteFile,
				RemotePath: newRemotePath,
			})
		}
	}

	return nil
}

// Recursively scans a folder to delete all files it contains.
func needDeleteDir(remoteDirPath string, client remote_clients.IClient) error {
	handle := func(err error) error {
		return err
	}

	remoteFiles, err := client.ReadDir(remoteDirPath)
	if err != nil {
		return handle(err)
	}

	// Because the queue is a stack you must add the delete folder before deleting the files,
	// this ensures that the files are deleted before the folder.
	taskQueue.AddTask(&tasks.Task{
		TypeTask:   tasks.DeleteFile,
		RemotePath: remoteDirPath,
	})

	// Recursively delete files from folder.
	for _, remoteFile := range remoteFiles {
		switch {
		case remoteFile.IsDir():
			// Call to continue scanning files in the new-found folder.

			newRemoteDirPath := path.Join(remoteDirPath, remoteFile.Name())

			err := needDeleteDir(newRemoteDirPath, client)
			if err != nil {
				return handle(err)
			}

		case !remoteFile.IsDir():
			// Call when file is found.

			newRemoteFilePath := path.Join(remoteDirPath, remoteFile.Name())

			taskQueue.AddTask(&tasks.Task{
				TypeTask:   tasks.DeleteFile,
				RemotePath: newRemoteFilePath,
			})
		}
	}

	return nil
}

// Recursively scans a folder to upload all files it contains.
func needUploadDir(localDirPath string, remoteDirPath string) error {
	handle := func(err error) error {
		return err
	}

	localFiles, err := extractFileInfoOfDir(localDirPath)
	if err != nil {
		return handle(err)
	}

	// Recursively upload files from folder.
	for _, localFile := range localFiles {
		switch {
		case localFile.IsDir():
			// Call to continue scanning files in the new-found folder.

			newLocalDirPath := path.Join(localDirPath, localFile.Name())
			newRemoteDirPath := path.Join(remoteDirPath, utils.NormalizedFileName(localFile.Name()))

			err := needUploadDir(newLocalDirPath, newRemoteDirPath)
			if err != nil {
				return handle(err)
			}

		case !localFile.IsDir():
			// Call when file is found.

			newLocalFilePath := path.Join(localDirPath, localFile.Name())
			newRemoteFilePath := path.Join(remoteDirPath, utils.NormalizedFileName(localFile.Name()))

			taskQueue.AddTask(&tasks.Task{
				TypeTask:   tasks.UploadFile,
				LocalPath:  newLocalFilePath,
				RemotePath: newRemoteFilePath,
				NeedUpload: localFile.Size(),
			})
		}
	}

	// Because the queue is a stack you must add the creation folder after uploading files,
	// this ensures that the folder are created before the uploading the files.
	taskQueue.AddTask(&tasks.Task{
		TypeTask:   tasks.CreateDir,
		LocalPath:  localDirPath,
		RemotePath: remoteDirPath,
	})

	return nil
}
