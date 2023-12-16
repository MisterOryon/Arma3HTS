package cmd

import (
	"fmt"
	"github.com/MisterOryon/Arma3HTS/mods"
	"github.com/MisterOryon/Arma3HTS/remote_clients"
	"github.com/MisterOryon/Arma3HTS/utils"
	"github.com/spf13/cobra"
	"os"
)

var checkRemoteConnexionCmd = &cobra.Command{
	Use:   "remote-connexion [preset path] [URI (scheme://user:password@host:port)]",
	Short: "Checks if all mods folder are present on your remote server",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		modList, err := mods.LoadMods(args[0])
		if err != nil {
			mods.DisplayLoadModError(err, debug)
			os.Exit(1)
		}

		remoteLogin, err := utils.ExtractIdentifiersOfUri(args[1])
		if err != nil {
			utils.DisplayIdentifierExtractorError(err, debug)
			os.Exit(1)
		}

		builder, err := remote_clients.NewBuilder(remoteLogin.Scheme)
		if err != nil {
			utils.DisplayNewRemoteBuilderError(err, debug, remoteLogin.Scheme)
			os.Exit(1)
		}

		builder.SetHost(remoteLogin.Host)
		builder.SetPort(remoteLogin.Port)
		builder.SetUser(remoteLogin.Username)
		builder.SetPassword(remoteLogin.Password)

		remoteClient, err := builder.GetClient()
		if err != nil {
			utils.DisplayBuildRemoteClientError(err, debug, remoteLogin.Scheme)
			os.Exit(1)
		}
		defer remoteClient.Close()

		errList := mods.CheckRemoteHostMods(remoteLogin.Path, &modList, remoteClient)
		if errList != nil {
			mods.DisplayCheckModErrors(errList, debug)
			os.Exit(1)
		}

		fmt.Println("successful")
	},
}
