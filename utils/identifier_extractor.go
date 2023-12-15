package utils

import (
	"net/url"
	"strconv"
	"strings"
)

// Identifier is a user-friendly struct for identifiers.
type Identifier struct {
	Scheme   string
	Host     string
	Port     int
	Username string
	Password string
	Path     string
}

// ExtractIdentifiersOfUri extract the identifiers of a URI to store them in Identifier struct more user-friendly.
func ExtractIdentifiersOfUri(uri string) (*Identifier, error) {
	handle := func(err error) (*Identifier, error) {
		return nil, err
	}

	parsedUrl, err := url.Parse(uri)
	if err != nil {
		return handle(err)
	}

	var (
		host string
		port int
	)

	scheme := parsedUrl.Scheme
	username := parsedUrl.User.Username()
	password, _ := parsedUrl.User.Password()
	path := parsedUrl.Path

	if hostFull := parsedUrl.Host; strings.Contains(hostFull, ":") {
		hostSplit := strings.Split(hostFull, ":")

		host = hostSplit[0]

		port, err = strconv.Atoi(hostSplit[1])
		if err != nil {
			return handle(err)
		}
	} else {
		host = hostFull
	}

	return &Identifier{
		Scheme:   scheme,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Path:     path,
	}, err
}
