package cmd

import (
	"fmt"
	"github.com/nicokempe/dnskit/pkg/version"
	"github.com/spf13/cobra"
)

var (
	// outputJSON emits JSON instead of colorized text.
	outputJSON bool

	// resolver overrides the system DNS resolver (host[:port]).
	resolver string
)

// rootCmd is the base command for DNSKit.
var rootCmd = &cobra.Command{
	Use:   "dnskit",
	Short: "DNSKit - A modern DNS analysis CLI",
	Long: `DNSKit is a modern and structured command-line tool
designed for DNS analysis, penetration testing, and system administration.`,
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DNSKit! Type --help or -h to see available commands.")
		fmt.Println("DNSKit is a modern and structured command-line tool designed for DNS analysis, penetration testing, and system administration.")
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&outputJSON, "json", false, "Output results in JSON format")
	rootCmd.PersistentFlags().StringVar(&resolver, "resolver", "", "Use a custom DNS resolver (ip[:port])")
}
