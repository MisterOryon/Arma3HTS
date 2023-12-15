package mods

import (
	"fmt"
)

// DisplayModList display a list of mods in the terminal.
func DisplayModList(modList []mod, listName bool, listId bool) {
	switch {
	case listName && listId:
		fmt.Printf("successful, %d mods found:\n", len(modList))

		for _, val := range modList {
			fmt.Printf("  - %s, @%d\n", val.GetName(), val.GetId())
		}

	case listName:
		fmt.Printf("successful, %d mods found:\n", len(modList))

		for _, val := range modList {
			fmt.Printf("  - %s\n", val.GetName())
		}

	case listId:
		fmt.Printf("successful, %d mods found: ", len(modList))

		for i, val := range modList {
			fmt.Printf("@%d", val.GetId())

			if i != len(modList)-1 {
				fmt.Print(";")
			}
		}

		fmt.Printf("\n")

	default:
		fmt.Printf("successful, %d mods found.\n", len(modList))
	}
}
