/*
Copyright © 2024 MARIA NIRANJAN <s19835@sci.pdn.ac.lk>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ip-traker-go",
	Short: "Trace the geographical location and other details of an IP address.",
	Long:  `This CLI application allows users to trace the geographical location, ISP, and other details of a given IP address. It provides essential information such as the country, city, and latitude/longitude coordinates. Additionally, the tool can be used to analyze multiple IP addresses at once, making it a useful utility for network administrators, cybersecurity analysts, and general users who want to gain insights into IP address origins. It supports both IPv4 and IPv6 formats, and can also perform reverse lookups if necessary.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ip-traker-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
