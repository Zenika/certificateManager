// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/configCommands.go

package cmd

import (
	"cm/config"
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration file management",
	Long:  `This is where you can create a templated file, edit/delete an existing config file, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var configEditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"update"},
	Short:   "Edit a configuration file",
	//Long:  `This is where you can create a templated file, edit/delete an existing config file, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.EditConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

var configTemplateCmd = &cobra.Command{
	Use: "template",
	//Aliases: []string{"update"},
	Short: "Create a templagte (blank) file",
	//Long:  `This is where you can create a templated file, edit/delete an existing config file, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		//err := config.TemplateConfigCreate()
		//if err != nil {
		//	fmt.Println(err)
		//}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configEditCmd)
	//configCmd.AddCommand(configTemplateCmd)
}
