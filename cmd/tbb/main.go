package main

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

const flagDataDir = "datadir"

func main() {
	var tbbCmd = &cobra.Command{
		Use: "tbb",
		Short: "The Blockchain bar",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	tbbCmd.AddCommand(versionCmd)
	tbbCmd.AddCommand(balancesCmd())
	tbbCmd.AddCommand(txCmd())
	tbbCmd.AddCommand(runCmd())
	
	err := tbbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}

}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}

func addDefaultRequiredFlags(cmd*cobra.Command) {
	cmd.Flags().String(flagDataDir, "","Absolute path to the node of directory")
	cmd.MarkFlagRequired(flagDataDir)
}