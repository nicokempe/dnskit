package cmd

import (
	"fmt"

	"github.com/nicokempe/dnskit/pkg/version"
	"github.com/spf13/cobra"
)

// versionCmd prints the current DNSKit version.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Info())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
