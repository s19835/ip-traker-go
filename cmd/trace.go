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
		fmt.Println("Testing race sub-command")
	},
}

func init() {
	rootCmd.AddCommand(trackeCmd)
}
