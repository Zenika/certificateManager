/*
©2023 J.F.Gratton (jean-francois@famillegratton.net)
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var version = "0.100 (2023.03.20)"

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
	//cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().BoolVarP(&misc.PlainOutput, "plain", "P", false, "Tables are shown with less decorations")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//func initConfig() {
//	//hosts.ConnectURI, _ = rootCmd.Flags().GetString("remotehost")
//	if !strings.Contains(hosts.ConnectURI, ":") {
//		hosts.ConnectURI += ":2375"
//	}
//}
