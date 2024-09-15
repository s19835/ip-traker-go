package root

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ip",
		Short: "ip-tracker-go",
		Long:  "ip address tracker CLI application using go with go modules",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
