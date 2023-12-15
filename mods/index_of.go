package mods

import (
	"github.com/MisterOryon/Arma3HTS/utils"
	"os"
)

// IndexOfFileInfoArray searches a file name in a list of FileInfo and returns its index or -1 one if it was not found.
func IndexOfFileInfoArray(array *[]os.FileInfo, fileName string) int {
	fileName = utils.NormalizedFileName(fileName)

	for index, arrayValue := range *array {
		if arrayValue.Name() == fileName {
			return index
		}
	}

	return -1
}

// IndexOfModsArray searches a mods id in a list of mods and returns its index or -1 one if it was not found.
func IndexOfModsArray(array *[]mod, idMod int) int {
	for index, arrayValue := range *array {
		if arrayValue.GetId() == idMod {
			return index
		}
	}

	return -1
}
