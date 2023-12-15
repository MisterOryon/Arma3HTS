package cmd

import (
	"fmt"
	"github.com/MisterOryon/Arma3HTS/utils"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Arma3HTS",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := utils.FindMainModuleVersion()

		switch {
		case err != nil && debug:
			fmt.Printf("Err: %+v\n", err)
			os.Exit(1)

		case err != nil:
			fmt.Println("Unable to get version number")
			os.Exit(1)

		default:
			fmt.Println(version)
		}
	},
}
