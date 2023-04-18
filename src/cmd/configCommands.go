// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/configCommands.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the configs command
var configCmd = &cobra.Command{
	Use:   "configs",
	Short: "Configuration file management",
	Long:  `This is where you can create a templated file, edit/delete an existing config file, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configs called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
