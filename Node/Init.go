package main

import (
	"github.com/spf13/cobra"
)

var option string
var rootCmd = &cobra.Command{
	Use:   "address",
	Short: "Parse address to start node",
	Run: func(cmd *cobra.Command, args []string) {
		SetAddress(option)
	},
}

//var setCmd = &cobra.Command{
//	Use: "address",
//	Run: func(cmd *cobra.Command, args []string) {
//		SetAddress(option)
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(setCmd)
//}

func Execute() {
	//rootCmd.Flags().String(option, "", "Set address")
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
