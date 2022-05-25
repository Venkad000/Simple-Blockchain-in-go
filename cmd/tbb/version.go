package main

import (
	"github.com/spf13/cobra"
	"fmt"
)

const Major = "0"
const Minor = "1"
const Fix = "0"
const Verbal = "TX Add && Balances List"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Describes the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version : %s.%s.%s-beta %s\n",Major,Minor,Fix,Verbal)
	},
}