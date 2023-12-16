package cmd

import (
	"fmt"
	"github.com/MisterOryon/Arma3HTS/mods"
	"github.com/MisterOryon/Arma3HTS/remote_clients"
	"github.com/MisterOryon/Arma3HTS/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync-mods [preset path] [!Workshop path] [URI (scheme://user:password@host:port)]",
	Short: "Allows you to synchronise server mods with a preset files and your local mods",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		modList, err := mods.LoadMods(args[0])
		if err != nil {
			mods.DisplayLoadModError(err, debug)
			os.Exit(1)
		}

		remoteLogin, err := utils.ExtractIdentifiersOfUri(args[2])
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

		taskQueue, errList := mods.SyncMods(args[1], remoteLogin.Path, &modList, remoteClient)
		if errList != nil {
			mods.DisplaySyncModError(err, debug)
			os.Exit(1)
		}

		totalTasks := taskQueue.Len()
		totalNeedUpload := taskQueue.NeedUpload()

		fmt.Printf("%d tasks are waiting and %s will be sent to the serve,\n", totalTasks, utils.ConvertSizeToStr(totalNeedUpload))
		fmt.Print("do you want to apply the changes? (y/n): ")

		var apply string
		_, err = fmt.Scanln(&apply)
		if err != nil {
			utils.DisplayScanlnError(err, debug)
			os.Exit(1)
		}

		if strings.TrimSpace(apply) != "y" {
			return
		}

		for taskQueue.AsNext() {
			currantTask := taskQueue.Next()
			taskNumber := totalTasks - taskQueue.Len()

			fmt.Printf("tasks (%d/%d)\n", taskNumber, totalTasks)

			err := currantTask.Run(remoteClient)
			if err != nil {
				utils.DisplayRunTaskError(err, debug)
				os.Exit(1)
			}
		}

		fmt.Println("successful")
	},
}
