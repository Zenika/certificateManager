// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/caCommands.go

package cmd

import (
	"certificateManager/ca"
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

var caCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "rootCA creation command",
	//Long:    `This is where you will manage (add/remove) your rootCAs\' configs files.`,
	Run: func(cmd *cobra.Command, args []string) {
		ca.CreateRootCA()
	},
}

var caVerifyCmd = &cobra.Command{
	Use:   "verify certificate_filename",
	Short: "verify the created CA certificate",
	//Long:    `This is where you will manage (add/remove) your rootCAs\' configs files.`,
	Run: func(cmd *cobra.Command, args []string) {
		ca.VerifyCACertificate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(caCmd)
	caCmd.AddCommand(caConfigCmd)
	caCmd.AddCommand(caCreateCmd)
	caCmd.AddCommand(caVerifyCmd)

	caVerifyCmd.Flags().BoolVarP(&ca.CaVerifyVerbose, "verbose", "v", false, "Display the full output")
}
