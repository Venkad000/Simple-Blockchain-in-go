package main

import (
	"fmt"
	"os"
	"blockchain/database"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use: "Balances",
		Short: "Interact with balances",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd* cobra.Command, args []string) {

		},
	}
	balancesCmd.AddCommand(balancesListCmd)
	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use: "List",
	Short: "List all balances",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}

		defer state.Close()

		fmt.Println("Accounts balances")
		fmt.Println("_________________")
		for account, balances:= range state.Balances {
			fmt.Printf("%s: %d\n",account, balances)
		}
	},
}

