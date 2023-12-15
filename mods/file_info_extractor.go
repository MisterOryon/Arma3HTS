package mods

import "os"

// extractFileInfoOfDir the os.ReadDir function returns []os.DirEntry but for the analysis of modes we need a list of []os.FileInfo,
// this function therefore extracts this list from os.ReadDir function.
func extractFileInfoOfDir(dirPath string) ([]os.FileInfo, error) {
	handle := func(err error) ([]os.FileInfo, error) {
		return nil, err
	}

	dirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return handle(err)
	}

	var filesInfo []os.FileInfo

	for _, currantFile := range dirEntry {
		currantFileInfo, err := currantFile.Info()
		if err != nil {
			return handle(err)
		}

		filesInfo = append(filesInfo, currantFileInfo)
	}

	return filesInfo, nil
}
