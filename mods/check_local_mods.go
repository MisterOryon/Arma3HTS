package mods

import (
	"errors"
	"os"
	"path"
)

// CheckLocalHostMods allows you to check if the !workshop folder contains all mods required by the preset.
func CheckLocalHostMods(workshopPath string, listMods *[]mod) []error {
	var modErrors []error

	for _, mod := range *listMods {
		modName := "@" + mod.GetName()
		modPath := path.Join(workshopPath, modName)

		localFile, err := os.Stat(modPath)

		switch {
		case os.IsNotExist(err):
			modErrors = append(modErrors, &CheckModError{
				_type: DirNotFound,
				path:  modPath,
				err:   err,
			})

		case err != nil:
			modErrors = append(modErrors, &CheckModError{
				_type: Unknown,
				path:  modPath,
				err:   err,
			})

		case !localFile.IsDir():
			modErrors = append(modErrors, &CheckModError{
				_type: IsNotDir,
				path:  modPath,
				err:   errors.New(CheckModErrorType.String(IsNotDir)),
			})
		}
	}

	return modErrors
}
