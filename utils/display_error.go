package utils

import "fmt"

// DisplayNewRemoteBuilderError display a new remote builder error in the terminal.
func DisplayNewRemoteBuilderError(err error, debug bool, remoteClientType string) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Printf("The %s remote client is not available!\n", remoteClientType)
	}
}

// DisplayBuildRemoteClientError display a build remote client error in the terminal.
func DisplayBuildRemoteClientError(err error, debug bool, remoteClientType string) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Printf("An error occurred while creating the %s client!\n", remoteClientType)
	}
}

// DisplayIdentifierExtractorError display an identifier extractor error in the terminal.
func DisplayIdentifierExtractorError(err error, debug bool) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Println("Unable to retrieve credentials from URI, are you sure this is correct?")
	}
}

// DisplayScanlnError display a scanln error in the terminal.
func DisplayScanlnError(err error, debug bool) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Println("Unable to read terminal input!")
	}
}

// DisplayRunTaskError display a run task error in the terminal.
func DisplayRunTaskError(err error, debug bool) {
	switch {
	case err != nil && debug:
		fmt.Printf("Err: %+v\n", err)

	default:
		fmt.Println("Unable to run task!")
	}
}
