package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"blockchain/node"
)



func runCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use: "run",
		Short: "Launces the TBB node and its HTTP api",
		Run: func(cmd *cobra.Command,args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)
			port, _ := cmd.Flags().GetUint64(flagPort)

			fmt.Printf("Launching the TBB node and its HTTP API...\n\n")

			bootstrap := node.NewPeerNode(
				"127.0.0.1",
				8080,
				true,
				true,
			)

			n := node.New(dataDir, port,bootstrap)
			err := n.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		},
	}

	addDefaultRequiredFlags(runCmd)
	runCmd.Flags().Uint64(flagPort,node.DefaultHttpPort,"Deafult port for communication of peers")

	return runCmd
}