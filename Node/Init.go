package main

import "github.com/spf13/cobra"

var option string
var rootCmd = &cobra.Command{
	Use:   "address",
	Short: "Parse address to start node",
	Run: func(cmd *cobra.Command, args []string) {
		SetAddress(option)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
