package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var trackeCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace IP address in a network",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Println("Provide valid IP address")
		}
	},
}

func init() {
	rootCmd.AddCommand(trackeCmd)
}
