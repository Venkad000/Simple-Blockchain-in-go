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

			fmt.Println("Launching the TBB node and its HTTP API...")

			err := node.Run(dataDir)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		},
	}

	addDefaultRequiredFlags(runCmd)

	return runCmd
}