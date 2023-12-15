package cmd

import (
	"fmt"
	"github.com/MisterOryon/Arma3HTS/mods"
	"github.com/spf13/cobra"
	"os"
)

var checkLocalModsCmd = &cobra.Command{
	Use:   "local-mods [preset path] [!Workshop path]",
	Short: "Check if all mods folders are present on the machine",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		modList, err := mods.LoadMods(args[0])
		if err != nil {
			mods.DisplayLoadModError(err, debug)
			os.Exit(1)
		}

		errList := mods.CheckLocalHostMods(args[1], &modList)
		if errList != nil {
			mods.DisplayCheckModErrors(errList, debug)
			os.Exit(1)
		}

		fmt.Println("successful")
	},
}
