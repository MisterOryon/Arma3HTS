package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "Arma3HTS"}

var debug bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "verbose debug output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
