// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/caCommands.go

package cmd

import (
	"certificateManager/ca"
	"certificateManager/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var privKeySize int

// caCmd represents the ca command
var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Root Certificate Authority management",
	Long:  `This is where you will manage (add/verify) your rootCAs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage ca {create|verify")
			os.Exit(0)
		}
	},
}

var caCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "rootCA creation command",
	//Long:    `This is where you will manage (add/remove) your rootCAs\' config files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ca.CreateRootCA(privKeySize)
		if err != nil {
			fmt.Println("Error while creating the root CA:")
			fmt.Println(err)
		} else {
			fmt.Printf("Certificate %s with a keysize of %v bits has been created in %s\n", config.CertConfig.CertificateName, privKeySize, config.CertConfig.CertificateDirectory)
			//fmt.Println("Certificate has been created.")
		}
	},
}

var caVerifyCmd = &cobra.Command{
	Use:   "verify certificate_filename",
	Short: "verify the created CA certificate",
	//Long:    `This is where you will manage (add/remove) your rootCAs\' config files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to provide the certificate name to be verified")
			os.Exit(0)
		}
		err := ca.VerifyCACertificate(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(caCmd)
	caCmd.AddCommand(caCreateCmd)
	caCmd.AddCommand(caVerifyCmd)

	caVerifyCmd.Flags().BoolVarP(&ca.CaVerifyVerbose, "verbose", "v", false, "Display the full output")
	caCreateCmd.Flags().IntVarP(&privKeySize, "keysize", "b", 4096, "Certificate private key size in bits")
}
