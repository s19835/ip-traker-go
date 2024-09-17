/*
Copyright © 2024 MARIA NIRANJAN <s19835@sci.pdn.ac.lk>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP address details using the ipinfo API.",
	Long:  `The "trace" command allows you to trace and retrieve detailed information about an IP address using the ipinfo API. This command provides key insights such as the IP's geographical location, including country, city, and coordinates, as well as the Internet Service Provider (ISP) and organization details. It supports both IPv4 and IPv6 addresses. By leveraging the ipinfo API, the "trace" command ensures accurate and up-to-date information, making it a powerful tool for network diagnostics, security analysis, and general IP tracking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trace called")
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
