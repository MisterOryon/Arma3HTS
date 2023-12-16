package cmd

import (
	"github.com/spf13/cobra"
)

var (
	listName bool
	listId   bool
)

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.AddCommand(checkPresetCmd)
	checkCmd.AddCommand(checkLocalModsCmd)
	checkCmd.AddCommand(checkRemoteConnexionCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Allows you to check a preset files, locale mods and remote mods",
}
