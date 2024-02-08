package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rum",
	Short: "Application to make pull requests easier",
	Long:  "Rum is an application that helps you to automate your pull requests.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	viewPrCommand.Flags().StringVar(&viewPrStatus, "status", "OPEN", "the status of pull requests to view")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
