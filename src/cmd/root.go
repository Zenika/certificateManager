/*
©2023 J.F.Gratton (jean-francois@famillegratton.net)
*/
package cmd

import (
	"certificateManager/configs"
	"certificateManager/misc"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.101 (2023.03.20)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "certificateManager {ca|config|cert}",
	Short:   "A rootCA and server certificates management tool",
	Version: version,
	Long:    `This tools allows you to manipulate your custom root CAs and all certificates signed against that rootCA.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		misc.Changelog()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	//fmt.Println("Remote Host is ==>", hosts.ConnectURI)
}

func init() {
	rootCmd.AddCommand(clCmd)
	rootCmd.PersistentFlags().StringVarP(&configs.CertConfigFile, "config", "c", "rootCA-default.json", "Root CA configuration file.")
}

//func initConfig() {
//	//hosts.ConnectURI, _ = rootCmd.Flags().GetString("remotehost")
//	if !strings.Contains(hosts.ConnectURI, ":") {
//		hosts.ConnectURI += ":2375"
//	}
//}
