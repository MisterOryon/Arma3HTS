package mods

import (
	"errors"
	"fmt"
)

// DisplayCheckModErrors display a check mod errors in the terminal.
func DisplayCheckModErrors(errList []error, debug bool) {
	fmt.Println("The check failed, the following error(s) were encountered:")

	for _, err := range errList {
		var checkModError *CheckModError

		if ok := errors.As(err, &checkModError); debug && ok {
			fmt.Println("  -", checkModError.GetErr())
		} else {
			fmt.Println("  -", err.Error())
		}
	}
}

// DisplayLoadModError display a load mod error in the terminal.
func DisplayLoadModError(err error, debug bool) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Println("Unable to read the file were you sure the file is an Arma3 preset?")
	}
}

// DisplaySyncModError display a sync mod error in the terminal.
func DisplaySyncModError(err error, debug bool) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Println("An error occurred while parsing the differences between server modes and those on your local machine!")
	}
}
