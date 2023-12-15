package mods

import (
	"errors"
	"github.com/MisterOryon/Arma3HTS/remote_clients"
	"path"
	"strconv"
	"strings"
)

// CheckRemoteHostMods allows you to check if the server folder contains all mods folder required by the preset.
func CheckRemoteHostMods(remotePath string, listMods *[]mod, client remote_clients.IClient) []error {
	handle := func(err []error) []error {
		return err
	}

	var modErrors []error

	remoteDir, err := client.ReadDir(remotePath)
	if err != nil {
		return handle([]error{err})
	}

	// Scans remote client folders to find mod folders.
	for _, mod := range remoteDir {
		// Don't look folders which are not mod folders.
		if !strings.HasPrefix(mod.Name(), "@") {
			continue
		}

		modPath := path.Join(remotePath, mod.Name())

		// Delete the @ at the beginning of the mod name.
		modName, err := strconv.Atoi(mod.Name()[1:])
		if err != nil {
			return handle([]error{&CheckModError{
				_type: Unknown,
				path:  modPath,
				err:   err,
			}})
		}

		index := IndexOfModsArray(listMods, modName)

		switch {
		case index == -1:
			// If the folder was not found in the list of mods it is because it is an additional folder,
			// which must be deleted because it is no longer used by the preset.

			modErrors = append(modErrors, &CheckModError{
				_type: DirUseless,
				path:  modPath,
				err:   errors.New(DirUseless.String()),
			})

		case !mod.IsDir():
			modErrors = append(modErrors, &CheckModError{
				_type: IsNotDir,
				path:  modPath,
				err:   errors.New(IsNotDir.String()),
			})

			// Utils to prevent an DirNotFound error from being generated in addition to IsNotDir.
			fallthrough

		default:
			*listMods = append((*listMods)[:index], (*listMods)[index+1:]...)
		}
	}

	// After analyzing of folder, only not found folders remain in listMods since they are not present on the remote client,
	// it must be necessary to uploading from the server.
	for _, modsNotFound := range *listMods {
		modPath := path.Join(remotePath, "@"+strconv.Itoa(modsNotFound.GetId()))

		modErrors = append(modErrors, &CheckModError{
			_type: DirNotFound,
			path:  modPath,
			err:   errors.New(DirNotFound.String()),
		})
	}

	return modErrors
}
