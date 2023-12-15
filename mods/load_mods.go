package mods

import (
	"errors"
	"golang.org/x/net/html"
	"os"
	"strconv"
	"strings"
)

var errFailedReadModContainer = errors.New("unable to read ModContainer tag")

// LoadMods returns the list of mods in the html preset.
func LoadMods(filePath string) ([]mod, error) {
	handle := func(err error) ([]mod, error) {
		return nil, err
	}

	bs, err := os.ReadFile(filePath)
	if err != nil {
		return handle(err)
	}

	modList, err := parse(string(bs))
	if err != nil {
		return handle(err)
	}

	return modList, nil
}

// Parse html preset.
func parse(text string) (modList []mod, err error) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var isModContainer bool
	var isDisplayName bool

	var displayName string
	var id int

	for {
		if err != nil {
			return
		}

		next := tkn.Next()
		token := tkn.Token()

		switch {
		case next == html.ErrorToken:
			// Is called when has reached the end of the file.

			return

		case isDisplayName:
			// Read content in DisplayName tag.

			displayName = token.Data
			isDisplayName = false

		case isModContainer:
			// Read content in ModContainer tag.

			switch {
			case token.Data == "td":
				// It is necessary to call next loop to get the value of DisplayName tag.

				if attr, isExist := getAttribute(token.Attr, "data-type"); isExist && attr == "DisplayName" {
					isDisplayName = true
				}

			case token.Data == "a":
				if attr, isExist := getAttribute(token.Attr, "href"); isExist {
					id, err = strconv.Atoi(strings.Split(attr, "?id=")[1])
				}
			}

			// Continue into next case to check if it has reached at the end of ModContainer tag.
			fallthrough

		case next == html.StartTagToken:
			// Read tags in html file.

			switch {
			case token.Data != "tr":
				continue

			case isModContainer:
				// Is called to check if ModContainer has been fully read.

				if displayName == "" || id == 0 {
					err = errFailedReadModContainer
				}

				modList = append(modList, mod{id: id, name: displayName})

				displayName = ""
				id = 0
				isModContainer = false

			default:
				// Is called when research ModContainer tag.

				if attr, isExist := getAttribute(token.Attr, "data-type"); isExist && attr == "ModContainer" {
					isModContainer = true
				}
			}
		}
	}
}

// Returns the value of the attribute in case it is set, and whether it is set.
func getAttribute(attributes []html.Attribute, key string) (string, bool) {
	for _, attr := range attributes {
		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}
