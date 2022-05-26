package main

import (
	"github.com/spf13/cobra"
	"fmt"
)

const Major = "0"
const Minor = "69"
const Fix = "0"
const Verbal = "Transparent database"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Describes the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version : %s.%s.%s-beta %s\n",Major,Minor,Fix,Verbal)
	},
}