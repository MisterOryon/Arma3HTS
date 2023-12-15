package cmd

import (
	"github.com/MisterOryon/Arma3HTS/mods"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	checkPresetCmd.Flags().BoolVar(&listName, "list-name", false, "show list of mod names")
	checkPresetCmd.Flags().BoolVarP(&listId, "list-id", "l", false, "show list of mod ids")
}

var checkPresetCmd = &cobra.Command{
	Use:   "perset [preset path]",
	Short: "Check an Arma3 preset in html format",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modList, err := mods.LoadMods(args[0])
		if err != nil {
			mods.DisplayLoadModError(err, debug)
			os.Exit(1)
		}

		mods.DisplayModList(modList, listName, listId)
	},
}
