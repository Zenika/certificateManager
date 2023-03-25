/*
Copyright © 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// caCmd represents the ca command
var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Root Certificate Authority maangement",
	Long:  `This is where you will manage (add/remove) your rootCAs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough parameters")
			os.Exit(0)
		}
	},
}

var caConfigCmd = &cobra.Command{
	Use:     "configs",
	Aliases: []string{"cfg", "conf"},
	Short:   "rootCA configs file management",
	Long:    `This is where you will manage (add/remove) your rootCAs\' configs files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough parameters")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(caCmd)
	caCmd.AddCommand(caConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// caCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// caCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
